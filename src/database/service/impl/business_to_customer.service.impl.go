package impl

import (
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/database/service"
)

type businessToCustomerRepository struct {
	db *gorm.DB
}

func (repository *businessToCustomerRepository) SaveBusinessToCustomerRequest(request *model.BusinessToCustomer) (*model.BusinessToCustomer, error) {
	result := repository.db.Save(&request)
	if result.Error != nil {
		return nil, result.Error
	}
	var createdRow *model.BusinessToCustomer
	repository.db.Model(&model.BusinessToCustomer{ID: request.ID}).First(&createdRow)
	return createdRow, nil
}

func (repository *businessToCustomerRepository) GetBusinessToCustomerRequestById(requestId string) *model.BusinessToCustomer {
	var result *model.BusinessToCustomer
	repository.db.Model(&model.BusinessToCustomer{}).Where("originator_conversation_id = ?", requestId).First(&result)
	return result
}

func (repository *businessToCustomerRepository) UpdateBusinessToCustomerRequest(request *model.BusinessToCustomer) (int64, error) {
	result := repository.db.Save(&request)
	return result.RowsAffected, result.Error
}

func NewBusinessToCustomerService(db *gorm.DB) service.BusinessToCustomerService {
	var serviceInstance service.BusinessToCustomerService
	serviceInstance = &businessToCustomerRepository{
		db: db,
	}
	return serviceInstance
}
