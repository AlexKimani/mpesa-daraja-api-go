package impl

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/rest/dtos/request"
	"time"
)

type DatabaseInitiatorService struct {
	db *gorm.DB
}

// NewDatabaseInitiatorService pass the database connection to the service
func NewDatabaseInitiatorService(db *gorm.DB) *DatabaseInitiatorService {
	return &DatabaseInitiatorService{
		db: db,
	}
}

func (dis *DatabaseInitiatorService) CreateInitiator(initiatorRequest request.InitiatorRequest) (*model.Initiator, error) {
	initiator := &model.Initiator{
		InitiatorName:       initiatorRequest.InitiatorName,
		InitiatorCredential: initiatorRequest.InitiatorCredential,
	}
	result := dis.db.Create(&initiator)
	if result.Error != nil {
		log.Errorf("Failed to insert record: %+v due to error: %+v", initiatorRequest, result.Error)
		return nil, result.Error
	}
	var createdInitiator model.Initiator
	result = dis.db.First(&createdInitiator, initiator.ID)
	if result.Error != nil {
		log.Errorf("Failed to Find record with ID: %+v due to error: %+v", initiator.ID, result.Error)
		return nil, result.Error
	}
	return &createdInitiator, nil
}

func (dis *DatabaseInitiatorService) GetAllInitiators() (*[]model.Initiator, error) {
	var initiators []model.Initiator
	result := dis.db.Find(&initiators)
	if result.Error != nil {
		log.Errorf("Failed to Find records due to error: %+v", result.Error)
		return nil, result.Error
	}
	return &initiators, nil
}

func (dis *DatabaseInitiatorService) GetInitiatorById(id int64) (*model.Initiator, error) {
	var initiator model.Initiator
	result := dis.db.First(&initiator, id)
	if result.Error != nil {
		log.Errorf("Failed to Find record with ID: %+v due to error: %+v", id, result.Error)
		return nil, result.Error
	}
	return &initiator, nil
}

func (dis *DatabaseInitiatorService) UpdateInitiator(id int64, initiatorRequest request.InitiatorRequest) (*model.Initiator, error) {
	existingInitiator, err := dis.GetInitiatorById(id)
	if err != nil {
		log.Errorf("Failed to get initiator by id: %+v due to error: %+v", id, err)
		return nil, err
	}
	updateInitiator := &model.Initiator{
		ID:                  existingInitiator.ID,
		InitiatorName:       initiatorRequest.InitiatorName,
		InitiatorCredential: initiatorRequest.InitiatorCredential,
		CreatedAt:           existingInitiator.CreatedAt,
		UpdatedAt:           time.Now(),
	}
	result := dis.db.Updates(updateInitiator)
	if result.Error != nil {
		log.Errorf("Failed to update record: %+v due to error: %+v", initiatorRequest, result.Error)
		return nil, result.Error
	}
	return updateInitiator, nil
}
