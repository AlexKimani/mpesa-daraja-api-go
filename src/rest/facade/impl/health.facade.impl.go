package impl

import (
	"mpesa-daraja-api-go/src/database/service"
	"mpesa-daraja-api-go/src/rest/dtos/response"
	"mpesa-daraja-api-go/src/rest/facade"
)

type healthFacade struct {
	healthService service.HealthService
}

func (impl *healthFacade) GetSystemStats() response.HealthStatusResponse {
	return impl.healthService.GetServerStatus()
}

func NewHealthFacade(healthService service.HealthService) facade.HealthFacade {
	return &healthFacade{
		healthService: healthService,
	}
}
