package impl

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/database/service"
	"mpesa-daraja-api-go/src/rest/controllers"
	"mpesa-daraja-api-go/src/rest/dtos/request"
	"mpesa-daraja-api-go/src/rest/dtos/response"
	"mpesa-daraja-api-go/src/rest/facade"
	"net/http"
)

type initiatorFacade struct {
	initiator service.InitiatorService
}

func NewInitiatorFacade(initiatorService service.InitiatorService) facade.InitiatorFacade {
	return &initiatorFacade{
		initiator: initiatorService,
	}
}

func (ifl *initiatorFacade) SaveInitiator(initiatorRequest *request.InitiatorRequest) (*response.ApiResponse, error) {
	initiator := &model.Initiator{
		InitiatorName:       initiatorRequest.InitiatorName,
		InitiatorCredential: initiatorRequest.InitiatorCredential,
	}
	initiator, err := ifl.initiator.SaveInitiator(initiator)
	if err != nil {
		var errorMessage = fmt.Sprintf("Failed to create Initiator, error: %+v", err)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse(http.StatusInternalServerError, "Failed", errorMessage)
		return errorBody, err
	}
	apiResponse := controllers.GetApiResponse(http.StatusOK, "0", "Success", initiator)
	return apiResponse, nil
}

func (ifl *initiatorFacade) GetAllInitiators() *response.ApiResponse {
	initiators := ifl.initiator.GetAllInitiators()
	if initiators == nil {
		var errorMessage = fmt.Sprintf("Failed to get Initiators")
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse(http.StatusOK, "Failed", errorMessage)
		return errorBody
	}
	apiResponse := controllers.GetApiResponse(http.StatusOK, "0", "Success", initiators)
	return apiResponse
}

func (ifl *initiatorFacade) GetInitiatorById(id int64) *response.ApiResponse {
	initiator := ifl.initiator.GetInitiatorById(id)
	if initiator == nil {
		var errorMessage = fmt.Sprintf("Failed to get Initiator by ID %+v", id)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse(http.StatusNotFound, "Failed", errorMessage)
		return errorBody
	}
	apiResponse := controllers.GetApiResponse(http.StatusOK, "0", "Success", initiator)
	return apiResponse
}

func (ifl *initiatorFacade) GetInitiatorByName(initiatorName string) *response.ApiResponse {
	initiator := ifl.initiator.GetInitiatorByName(initiatorName)
	if initiator == nil {
		var errorMessage = fmt.Sprintf("Failed to get Initiator by name %+v", initiatorName)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse(http.StatusNotFound, "Failed", errorMessage)
		return errorBody
	}
	apiResponse := controllers.GetApiResponse(http.StatusOK, "0", "Success", initiator)
	return apiResponse
}

func (ifl *initiatorFacade) UpdateInitiator(id int64, initiatorRequest *request.InitiatorRequest) (*response.ApiResponse, error) {
	existingInitiator := ifl.initiator.GetInitiatorById(id)
	if existingInitiator == nil {
		var errorMessage = fmt.Sprintf("Could not find Initiator by ID: %+v", id)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse(http.StatusNotFound, "Failed", errorMessage)
		return errorBody, nil
	}

	updateRequest := &model.Initiator{
		ID:                  id,
		InitiatorName:       initiatorRequest.InitiatorName,
		InitiatorCredential: initiatorRequest.InitiatorCredential,
		CreatedAt:           existingInitiator.CreatedAt,
	}
	_, err := ifl.initiator.UpdateInitiator(updateRequest)
	if err != nil {
		var errorMessage = fmt.Sprintf("Failed to Update Initiator by ID: %+v, error: %+v", id, err)
		log.Errorf(errorMessage)
		errorBody := controllers.GetErrorResponse(http.StatusInternalServerError, "Failed", errorMessage)
		return errorBody, err
	}

	apiResponse := controllers.GetApiResponse(http.StatusOK, "0", fmt.Sprintf("Successfully updated initiator: %d", id), nil)
	return apiResponse, nil
}
