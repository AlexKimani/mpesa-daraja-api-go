package client

import (
	"github.com/gin-gonic/gin"
	"mpesa-daraja-api-go/src/client/dtos/results"
	"net/http"
)

func getC2BResult(context *gin.Context) {
	var c2bResultRequest results.MpesaC2bResult
	err := context.ShouldBindJSON(&c2bResultRequest)
	if err != nil {
		c2bResponse := results.MpesaC2bValidationResponse{
			ResultCode: "01",
			ResultDesc: "Bad Request",
		}
		context.JSON(http.StatusBadRequest, c2bResponse)
		return
	}
}
