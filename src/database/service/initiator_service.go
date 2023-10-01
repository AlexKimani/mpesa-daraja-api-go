package service

import (
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/rest/dtos/request"
)

type InitiatorService interface {
	CreateInitiator(request request.InitiatorRequest) (*model.Initiator, error)

	GetAllInitiators() (*[]model.Initiator, error)

	GetInitiatorById(id int64) (*model.Initiator, error)

	UpdateInitiator(id int64, initiatorRequest request.InitiatorRequest) (*model.Initiator, error)
}
