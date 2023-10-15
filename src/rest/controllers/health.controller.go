package controllers

import (
	"github.com/gin-gonic/gin"
	"mpesa-daraja-api-go/src/rest/facade"
	"net/http"
)

type HealthController interface {
	GetServerStatus(ctx *gin.Context)
}

type healthController struct {
	healthStats facade.HealthFacade
}

// GetServerStatus godoc
// @Summary      Get Server Stats
// @Description  get Server stats like DB and connection pool stats
// @Tags         Stats
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.HealthStatusResponse
// @Failure      400  {object}  response.HealthStatusResponse
// @Failure      404  {object}  response.HealthStatusResponse
// @Failure      500  {object}  response.HealthStatusResponse
// @Router       /health [get]
func (controller healthController) GetServerStatus(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, controller.healthStats.GetSystemStats())
}

func NewHealthController(healthStats facade.HealthFacade) HealthController {
	controllerInstance := &healthController{
		healthStats: healthStats,
	}
	return controllerInstance
}
