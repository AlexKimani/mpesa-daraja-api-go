package service

import "mpesa-daraja-api-go/src/database/model"

type BusinessToCustomerService interface {
	SaveBusinessToCustomerRequest(request *model.BusinessToCustomer) (*model.BusinessToCustomer, error)
	GetBusinessToCustomerRequestById(requestId string) *model.BusinessToCustomer
	UpdateBusinessToCustomerRequest(request *model.BusinessToCustomer) (int64, error)
}
