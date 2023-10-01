package facade

import (
	"mpesa-daraja-api-go/src/rest/dtos/request"
	"mpesa-daraja-api-go/src/rest/dtos/response"
)

type InitiatorFacade interface {
	CreateInitiator(request request.InitiatorRequest) *response.ApiResponse

	GetAllInitiators() *response.ApiResponse

	GetInitiatorById(id int64) *response.ApiResponse

	UpdateInitiator(id int64, initiatorRequest request.InitiatorRequest) *response.ApiResponse
}
