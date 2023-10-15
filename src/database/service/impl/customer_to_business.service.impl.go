package impl

import (
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/database/service"
)

type customerToBusinessRepository struct {
	db *gorm.DB
}

func (repository *customerToBusinessRepository) SaveCustomerToBusinessRequest(request *model.CustomerToBusiness) (*model.CustomerToBusiness, error) {
	result := repository.db.Save(&request)
	if result.Error != nil {
		return nil, result.Error
	}
	var createdRow *model.CustomerToBusiness
	repository.db.Model(&model.CustomerToBusiness{ID: request.ID}).First(&createdRow)
	return createdRow, nil
}

func (repository *customerToBusinessRepository) GetCustomerToBusinessRequestById(requestId string) *model.CustomerToBusiness {
	var result *model.CustomerToBusiness
	repository.db.Model(&model.CustomerToBusiness{}).Where("transaction_id = ?", requestId).First(&result)
	return result
}

func (repository *customerToBusinessRepository) UpdateCustomerToBusinessRequest(request *model.CustomerToBusiness) (int64, error) {
	result := repository.db.Save(&request)
	return result.RowsAffected, result.Error
}

func NewCustomerToBusinessService(db *gorm.DB) service.CustomerToBusinessService {
	var serviceInstance service.CustomerToBusinessService
	serviceInstance = &customerToBusinessRepository{
		db: db,
	}
	return serviceInstance
}
