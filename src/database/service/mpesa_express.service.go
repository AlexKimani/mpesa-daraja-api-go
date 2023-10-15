package service

import (
	"mpesa-daraja-api-go/src/database/model"
)

type MpesaExpressService interface {
	SaveMpesaExpressRequest(request *model.MpesaExpress) (*model.MpesaExpress, error)
	GetMpesaExpressRequestById(requestId string) *model.MpesaExpress
	UpdateMpesaExpressRequest(request *model.MpesaExpress) (int64, error)
}
