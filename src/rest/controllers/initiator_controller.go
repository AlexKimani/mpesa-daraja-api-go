package controllers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"mpesa-daraja-api-go/src/config"
	"mpesa-daraja-api-go/src/rest/dtos/request"
	"mpesa-daraja-api-go/src/rest/facade"
	"net/http"
)

type InitiatorApp struct {
	InitiatorFacade facade.InitiatorFacade
}

func InitiatorController(router *http.ServeMux) {
	router.HandleFunc("/initiator", CreateInitiatorHandler)
}

func CreateInitiatorHandler(writer http.ResponseWriter, requestServlet *http.Request) {
	// Verify Http Method
	apiResponse := ValidateHttpPostMethod(requestServlet)
	if apiResponse != nil {
		WriteResponseWriter(http.StatusMethodNotAllowed, apiResponse, writer)
		return
	}
	// Read request body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("Unable to read request body: %+v", err)
			var errorMessage = fmt.Sprintf("Unable to read request body: %s", err)
			errorBody := GetErrorResponse("Failed", errorMessage)
			WriteResponseWriter(http.StatusInternalServerError, errorBody, writer)
			return
		}
	}(requestServlet.Body)

	requestBody, err := io.ReadAll(requestServlet.Body)
	if err != nil {
		log.Errorf("Unable to read request body: %+v", err)
		var errorMessage = fmt.Sprintf("Unable to read request body: %s", err)
		errorBody := GetErrorResponse("Failed", errorMessage)
		WriteResponseWriter(http.StatusInternalServerError, errorBody, writer)
		return
	}
	var initiatorRequest request.InitiatorRequest
	err = json.Unmarshal(requestBody, &initiatorRequest)
	if err != nil {
		log.Errorf("Unable to Unmarshal request body: %+v", err)
		var errorMessage = fmt.Sprintf("Unable to Unmarshal request body: %s", err)
		errorBody := GetErrorResponse("Failed", errorMessage)
		WriteResponseWriter(http.StatusInternalServerError, errorBody, writer)
		return
	}

	apiResponse = config.InitiatorApp.InitiatorFacade.CreateInitiator(initiatorRequest)
	WriteResponseWriter(http.StatusCreated, apiResponse, writer)
}
