package service

import (
	"mpesa-daraja-api-go/src/rest/dtos/response"
)

type HealthService interface {
	GetServerStatus() response.HealthStatusResponse
}
