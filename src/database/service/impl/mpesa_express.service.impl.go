package impl

import (
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/database/service"
)

type mpesaExpressRepository struct {
	db *gorm.DB
}

func (repository *mpesaExpressRepository) SaveMpesaExpressRequest(request *model.MpesaExpress) (*model.MpesaExpress, error) {
	result := repository.db.Save(&request)
	if result.Error != nil {
		return nil, result.Error
	}
	var createdRow *model.MpesaExpress
	repository.db.Model(&model.MpesaExpress{ID: request.ID}).First(&createdRow)
	return createdRow, nil
}

func (repository *mpesaExpressRepository) GetMpesaExpressRequestById(requestId string) *model.MpesaExpress {
	var result *model.MpesaExpress
	repository.db.Model(&model.MpesaExpress{}).Where("merchant_request_id", requestId).First(&result)
	return result
}

func (repository *mpesaExpressRepository) UpdateMpesaExpressRequest(request *model.MpesaExpress) (int64, error) {
	result := repository.db.Save(&request)
	return result.RowsAffected, result.Error
}

func NewMpesaExpressService(db *gorm.DB) service.MpesaExpressService {
	var serviceInstance service.MpesaExpressService
	serviceInstance = &mpesaExpressRepository{
		db: db,
	}
	return serviceInstance
}
