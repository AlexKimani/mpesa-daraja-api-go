package service

import "mpesa-daraja-api-go/src/database/model"

type CustomerToBusinessService interface {
	SaveCustomerToBusinessRequest(request *model.CustomerToBusiness) (*model.CustomerToBusiness, error)
	GetCustomerToBusinessRequestById(requestId string) *model.CustomerToBusiness
	UpdateCustomerToBusinessRequest(request *model.CustomerToBusiness) (int64, error)
}
