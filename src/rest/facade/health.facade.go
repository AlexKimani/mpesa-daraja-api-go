package facade

import "mpesa-daraja-api-go/src/rest/dtos/response"

type HealthFacade interface {
	GetSystemStats() response.HealthStatusResponse
}
