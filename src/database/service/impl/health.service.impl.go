package impl

import (
	"fmt"
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/service"
	"mpesa-daraja-api-go/src/rest/dtos/response"
)

type healthRepository struct {
	db *gorm.DB
}

func (repository healthRepository) GetServerStatus() response.HealthStatusResponse {
	var connectionStatus string
	dbCon, err := repository.db.DB()
	err = dbCon.Ping()
	if err != nil {
		connectionStatus = fmt.Sprintf("Connection error: %+v", err)
	} else {
		connectionStatus = "Connected!"
	}

	stats := response.Stats{
		MaxOpenConnections: dbCon.Stats().MaxOpenConnections,
		OpenConnections:    dbCon.Stats().OpenConnections,
		InUse:              dbCon.Stats().InUse,
		Idle:               dbCon.Stats().Idle,
	}

	return response.HealthStatusResponse{
		ServerStatus:   "Server status: UP!",
		DatabaseStatus: connectionStatus,
		DatabaseStats:  stats,
	}
}

func NewHealthService(db *gorm.DB) service.HealthService {
	var serviceInstance service.HealthService
	serviceInstance = &healthRepository{
		db: db,
	}
	return serviceInstance
}
