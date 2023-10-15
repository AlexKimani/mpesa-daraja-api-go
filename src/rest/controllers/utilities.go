package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mpesa-daraja-api-go/src/rest/dtos/response"
	"net/http"
	"strings"
	"time"
)

// GetErrorResponse format error response object
func GetErrorResponse(status int, errorCode string, errorMessage string) *response.ApiResponse {
	errorResponse := &response.ApiResponse{
		HttpStatus:   status,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		TimeStamp:    time.Now(),
	}
	return errorResponse
}

// GetApiResponse format api data response
func GetApiResponse(status int, responseCode string, responseMessage string, data any) *response.ApiResponse {
	apiResponse := &response.ApiResponse{
		HttpStatus:      status,
		TimeStamp:       time.Now(),
		Data:            data,
		ResponseCode:    responseCode,
		ResponseMessage: responseMessage,
	}
	return apiResponse
}

type responseErr struct {
	Field     string `json:"field"`
	Condition string `json:"condition"`
}

func RenderBindingErrors(ctx *gin.Context, validationError validator.ValidationErrors) {
	var responseErrs []responseErr
	for _, fieldError := range validationError {
		field := fieldError.Field()
		responseErrs = append(responseErrs, responseErr{
			Field:     strings.ToLower(field[:1]) + field[1:],
			Condition: fieldError.ActualTag(),
		})
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, responseErrs)
}
