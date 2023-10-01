package impl

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"mpesa-daraja-api-go/src/database/service"
	"mpesa-daraja-api-go/src/rest/controllers"
	"mpesa-daraja-api-go/src/rest/dtos/request"
	"mpesa-daraja-api-go/src/rest/dtos/response"
)

type InitiatorFacadeImpl struct {
	initiator service.InitiatorService
}

func ApiInitiatorFacade(initiatorService service.InitiatorService) *InitiatorFacadeImpl {
	return &InitiatorFacadeImpl{
		initiator: initiatorService,
	}
}

func (ifl *InitiatorFacadeImpl) CreateInitiator(initiatorRequest request.InitiatorRequest) *response.ApiResponse {
	initiator, err := ifl.initiator.CreateInitiator(initiatorRequest)
	if err != nil {
		var errorMessage = fmt.Sprintf("Failed to create Initiator, error: %+v", err)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse("Failed", errorMessage)
		return errorBody
	}
	apiResponse := controllers.GetApiResponse("0", "Success", initiator)
	return apiResponse
}

func (ifl *InitiatorFacadeImpl) GetAllInitiators() *response.ApiResponse {
	initiators, err := ifl.initiator.GetAllInitiators()
	if err != nil {
		var errorMessage = fmt.Sprintf("Failed to Get Initiators, error: %+v", err)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse("Failed", errorMessage)
		return errorBody
	}
	apiResponse := controllers.GetApiResponse("0", "Success", initiators)
	return apiResponse
}

func (ifl *InitiatorFacadeImpl) GetInitiatorById(id int64) *response.ApiResponse {
	initiator, err := ifl.initiator.GetInitiatorById(id)
	if err != nil {
		var errorMessage = fmt.Sprintf("Failed to Get Initiator by ID: %+v, error: %+v", id, err)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse("Failed", errorMessage)
		return errorBody
	}
	apiResponse := controllers.GetApiResponse("0", "Success", initiator)
	return apiResponse
}

func (ifl *InitiatorFacadeImpl) UpdateInitiator(id int64, initiatorRequest request.InitiatorRequest) *response.ApiResponse {
	initiator, err := ifl.initiator.UpdateInitiator(id, initiatorRequest)
	if err != nil {
		var errorMessage = fmt.Sprintf("Failed to Update Initiator by ID: %+v, error: %+v", id, err)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse("Failed", errorMessage)
		return errorBody
	}

	apiResponse := controllers.GetApiResponse("0", "Success", initiator)
	return apiResponse
}
