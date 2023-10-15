package config

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/service/impl"
	"mpesa-daraja-api-go/src/rest/controllers"
	impl2 "mpesa-daraja-api-go/src/rest/facade/impl"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitializeServices(engine *gin.Engine, db *gorm.DB) {
	// initialize Services
	initiatorService := impl.NewInitiatorService(db)
	healthService := impl.NewHealthService(db)

	// initialize facades
	healthFacade := impl2.NewHealthFacade(healthService)
	initiatorFacade := impl2.NewInitiatorFacade(initiatorService)

	// initialize controllers
	healthController := controllers.NewHealthController(healthFacade)
	initiatorController := controllers.NewInitiatorController(initiatorFacade)

	// initialize controller APIs
	v1 := engine.Group("/v1")
	{
		v1.GET("/health", healthController.GetServerStatus)
		v1.POST("/initiator/create", initiatorController.SaveInitiator)
		v1.GET("/initiators", initiatorController.GetAllInitiators)
		v1.GET("/initiator/id", initiatorController.GetInitiatorById)
		v1.GET("/initiator/name", initiatorController.GetInitiatorByName)
		v1.PATCH("/initiator/update", initiatorController.UpdateInitiator)
	}
}

func SwaggerSetup(config Config, engine *gin.Engine) {
	address := "http://" + config.Server.Host + ":" + config.Server.Port + "/swagger/doc.json"

	url := ginSwagger.URL(address)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// @title			Safaricom Daraja 2.0 Api
// @version			1.0
// @description		This service is emant to integrate with Safaricom Daraja 2.0 API
// @termsOfService	https://apache.org

// @contact.name	Joe Alex N Kimani
// @contact.url		https://wakatime.com/@joealexkimani
// @contact.email	joealex.kimani@gmail.com

// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html

// @host			0.0.0.0:8000
// @BasePath		/v1/

// Run will start the HTTP Server and initiate connection pool
func (configuration Config) Run() {
	// Set up a channel to listen to for interrupt signals
	var runChannel = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(configuration.Server.Timeout.Server),
	)
	defer cancel()
	// delay system start-up to allow mysql to initialize
	time.Sleep(time.Second * 5)

	// Connect to DB
	dbCon, err := ConnectToDatabase(configuration)

	// Call DB Connection pool start
	db, err := DatabaseConnectionPool(configuration, dbCon)
	if err != nil {
		log.Fatalf("Unable to initialize database connection pool due to error: %+v", err)
		// Send kill command to close server as DB is not available
		runChannel <- syscall.SIGINT
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Alert the user that the Database migrations are starting
	log.Info("Initializing Database Migrations")

	// Call DB Migrations
	err = RunDatabaseMigrations(db, configuration)
	if err != nil {
		log.Fatalf("Unable to initialize database migrations due to error: %+v", err)
		// Send kill command to close server as DB is not available
		runChannel <- syscall.SIGINT
	}

	if err != nil {
		log.Fatalf("Unable to initialize database connection pool due to error: %+v", err)
		// Send kill command to close server as DB is not available
		runChannel <- syscall.SIGINT
	}

	dbCon, _ = ConnectToDatabase(configuration)

	// Define server options
	router := gin.Default()

	//initialize services, facades and apis
	InitializeServices(router, dbCon)

	// initialize swagger docs
	SwaggerSetup(configuration, router)

	// Define server options
	server := &http.Server{
		Addr:         ":" + configuration.Server.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(configuration.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(configuration.Server.Timeout.Write) * time.Second,
		IdleTimeout:  time.Duration(configuration.Server.Timeout.Idle) * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Alert the user that the server is starting
	log.Infof("Server is starting on %s", server.Addr)

	// run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// Normal interrupt operation, ignore
			} else {
				log.Fatalf("Server failed to start due to error %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable, so we can let the user know why the server is shutting down
	interrupt := <-runChannel

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	log.Infof("Server is shutting down due to %+v", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
