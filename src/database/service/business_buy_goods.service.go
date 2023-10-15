package service

import "mpesa-daraja-api-go/src/database/model"

type BusinessBuyGoodsService interface {
	SaveBuyGoodsRequest(buyGoods *model.BusinessBuyGood) (*model.BusinessBuyGood, error)
	GetBuyGoodsTransactionById(requestId string) *model.BusinessBuyGood
	UpdateBuyGoodsRequest(buyGoods *model.BusinessBuyGood) (int64, error)
}
