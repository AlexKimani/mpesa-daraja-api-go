package impl

import (
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/database/service"
)

type initiatorRepository struct {
	db *gorm.DB
}

func NewInitiatorService(db *gorm.DB) service.InitiatorService {
	var serviceInstance service.InitiatorService
	serviceInstance = &initiatorRepository{
		db: db,
	}
	return serviceInstance
}

func (repository *initiatorRepository) SaveInitiator(initiator *model.Initiator) (*model.Initiator, error) {
	result := repository.db.Save(initiator)
	if result.Error != nil {
		return nil, result.Error
	}
	var createdRow *model.Initiator
	repository.db.Model(&model.Initiator{ID: initiator.ID}).First(&createdRow)
	return createdRow, nil
}

func (repository initiatorRepository) GetAllInitiators() *[]model.Initiator {
	var result *[]model.Initiator
	repository.db.Find(&result)
	return result
}

func (repository initiatorRepository) GetInitiatorById(id int64) *model.Initiator {
	var result *model.Initiator
	repository.db.Model(&model.Initiator{ID: id}).First(&result)
	return result
}

func (repository initiatorRepository) GetInitiatorByName(initiatorName string) *model.Initiator {
	var result *model.Initiator
	repository.db.Model(&model.Initiator{}).Where("initiator_name = ?", initiatorName).First(&result)
	return result
}

func (repository initiatorRepository) UpdateInitiator(initiator *model.Initiator) (int64, error) {
	result := repository.db.Save(&initiator)
	return result.RowsAffected, result.Error
}
