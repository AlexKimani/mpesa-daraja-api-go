package facade

import (
	"mpesa-daraja-api-go/src/rest/dtos/request"
	"mpesa-daraja-api-go/src/rest/dtos/response"
)

type InitiatorFacade interface {
	SaveInitiator(request *request.InitiatorRequest) (*response.ApiResponse, error)
	GetAllInitiators() *response.ApiResponse
	GetInitiatorById(id int64) *response.ApiResponse
	GetInitiatorByName(initiatorName string) *response.ApiResponse
	UpdateInitiator(id int64, initiatorRequest *request.InitiatorRequest) (*response.ApiResponse, error)
}
