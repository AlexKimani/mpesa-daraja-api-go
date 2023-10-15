package impl

import (
	"gorm.io/gorm"
	"mpesa-daraja-api-go/src/database/model"
	"mpesa-daraja-api-go/src/database/service"
)

type businessBuyGoodsRepository struct {
	db *gorm.DB
}

func (repository *businessBuyGoodsRepository) SaveBuyGoodsRequest(request *model.BusinessBuyGood) (*model.BusinessBuyGood, error) {
	result := repository.db.Save(&request)
	if result.Error != nil {
		return nil, result.Error
	}
	var createdRow *model.BusinessBuyGood
	repository.db.Model(&model.BusinessBuyGood{ID: request.ID}).First(&createdRow)
	return createdRow, nil
}

func (repository *businessBuyGoodsRepository) GetBuyGoodsTransactionById(requestId string) *model.BusinessBuyGood {
	var result *model.BusinessBuyGood
	repository.db.Model(&model.BusinessBuyGood{}).Where("originator_conversation_id = ?", requestId).First(&result)
	return result
}

func (repository *businessBuyGoodsRepository) UpdateBuyGoodsRequest(buyGoods *model.BusinessBuyGood) (int64, error) {
	result := repository.db.Save(&buyGoods)
	return result.RowsAffected, result.Error
}

func NewBusinessBuyGoodsService(db *gorm.DB) service.BusinessBuyGoodsService {
	var serviceInstance service.BusinessBuyGoodsService
	serviceInstance = &businessBuyGoodsRepository{
		db: db,
	}
	return serviceInstance
}
