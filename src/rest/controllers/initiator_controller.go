package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mpesa-daraja-api-go/src/rest/dtos/request"
	"mpesa-daraja-api-go/src/rest/facade"
	"net/http"
	"strconv"
)

type InitiatorController interface {
	SaveInitiator(ctx *gin.Context)
	GetAllInitiators(ctx *gin.Context)
	GetInitiatorById(ctx *gin.Context)
	GetInitiatorByName(ctx *gin.Context)
	UpdateInitiator(ctx *gin.Context)
}

type initiatorController struct {
	initiatorFacade facade.InitiatorFacade
}

// SaveInitiator godoc
// @Summary      Create an Initiator object
// @Description  Create initiator details to be used to API auth
// @Tags         initiator
// @Accept       json
// @Produce      json
// @Param        data body request.InitiatorRequest true  "initiator"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /initiators/create [post]
// Will need to fix the any datatype response
func (controller initiatorController) SaveInitiator(ctx *gin.Context) {
	requestModel := &request.InitiatorRequest{}
	if err := ctx.ShouldBind(requestModel); err != nil && errors.As(err, &validator.ValidationErrors{}) {
		RenderBindingErrors(ctx, err.(validator.ValidationErrors))
		return
	}
	response, err := controller.initiatorFacade.SaveInitiator(requestModel)
	if err != nil {
		ctx.IndentedJSON(response.HttpStatus, response)
		return
	}
	ctx.IndentedJSON(response.HttpStatus, response)
}

func (controller initiatorController) GetAllInitiators(ctx *gin.Context) {
	response := controller.initiatorFacade.GetAllInitiators()
	ctx.IndentedJSON(response.HttpStatus, response)
}

func (controller initiatorController) GetInitiatorById(ctx *gin.Context) {
	param := ctx.Query("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		errorResponse := GetErrorResponse(http.StatusBadRequest, "Bad Request", "&id param not provide")
		ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	apiResponse := controller.initiatorFacade.GetInitiatorById(int64(id))
	ctx.IndentedJSON(http.StatusOK, apiResponse)
}

func (controller initiatorController) GetInitiatorByName(ctx *gin.Context) {
	param := ctx.Query("name")
	if param == "" {
		errorResponse := GetErrorResponse(http.StatusBadRequest, "Bad Request", "name param not provide")
		ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	apiResponse := controller.initiatorFacade.GetInitiatorByName(param)
	ctx.IndentedJSON(http.StatusOK, apiResponse)
}

func (controller initiatorController) UpdateInitiator(ctx *gin.Context) {
	param := ctx.Query("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		errorResponse := GetErrorResponse(http.StatusBadRequest, "Bad Request", "id param not provide")
		ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	requestModel := &request.InitiatorRequest{}
	if err := ctx.ShouldBind(requestModel); err != nil && errors.As(err, &validator.ValidationErrors{}) {
		RenderBindingErrors(ctx, err.(validator.ValidationErrors))
		return
	}
	response, err := controller.initiatorFacade.UpdateInitiator(int64(id), requestModel)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, response)
		return
	}
	ctx.IndentedJSON(response.HttpStatus, response)
}

func NewInitiatorController(initiatorFacade facade.InitiatorFacade) InitiatorController {
	controllerInstance := &initiatorController{
		initiatorFacade: initiatorFacade,
	}
	return controllerInstance
}
