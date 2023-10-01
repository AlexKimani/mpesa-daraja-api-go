package controllers

import (
	"encoding/json"
	"fmt"
	"mpesa-daraja-api-go/src/rest/dtos/response"
	"net/http"
	"time"
)

func ValidateHttpPostMethod(request *http.Request) *response.ApiResponse {
	if request.Method != http.MethodPost {
		var errorMessage = fmt.Sprintf("Method %s Not Allowed.", request.Method)
		errorBody := GetErrorResponse("Not Allowed", errorMessage)
		return errorBody
	}
	return nil
}

func ValidateHttpGetMethod(request *http.Request) *response.ApiResponse {
	if request.Method != http.MethodGet {
		var errorMessage = fmt.Sprintf("Method %s Not Allowed.", request.Method)
		errorBody := GetErrorResponse("Not Allowed", errorMessage)
		return errorBody
	}
	return nil
}

func ValidateHttpPutMethod(request *http.Request) *response.ApiResponse {
	if request.Method != http.MethodPut {
		var errorMessage = fmt.Sprintf("Method %s Not Allowed.", request.Method)
		errorBody := GetErrorResponse("Not Allowed", errorMessage)
		return errorBody
	}
	return nil
}

// WriteResponseWriter writes the response body
func WriteResponseWriter(status int, responseBody any, writer http.ResponseWriter) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		return
	}
}

// GetErrorResponse format error response object
func GetErrorResponse(errorCode string, errorMessage string) *response.ApiResponse {
	errorResponse := &response.ApiResponse{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		TimeStamp:    time.Now(),
	}
	return errorResponse
}

// GetApiResponse format api data response
func GetApiResponse(responseCode string, responseMessage string, data any) *response.ApiResponse {
	apiResponse := &response.ApiResponse{
		TimeStamp:       time.Now(),
		Data:            data,
		ResponseCode:    responseCode,
		ResponseMessage: responseMessage,
	}
	return apiResponse
}
