package service

import (
	"mpesa-daraja-api-go/src/database/model"
)

type InitiatorService interface {
	SaveInitiator(initiator *model.Initiator) (*model.Initiator, error)
	GetAllInitiators() *[]model.Initiator
	GetInitiatorById(id int64) *model.Initiator
	GetInitiatorByName(initiatorName string) *model.Initiator
	UpdateInitiator(initiator *model.Initiator) (int64, error)
}
