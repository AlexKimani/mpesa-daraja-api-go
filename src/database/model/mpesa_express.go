package model

import (
	"time"
)

const TableNameMpesaExpress = "mpesa_express"

// MpesaExpress mapped from table <mpesa_express>
type MpesaExpress struct {
	ID                     int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BusinessShortCode      string    `gorm:"column:business_short_code;not null" json:"business_short_code"`
	Timestamp              string    `gorm:"column:timestamp" json:"timestamp"`
	TransactionType        string    `gorm:"column:transaction_type;not null" json:"transaction_type"`
	Amount                 float64   `gorm:"column:amount" json:"amount"`
	PartyA                 string    `gorm:"column:party_a;not null" json:"party_a"`
	PartyB                 string    `gorm:"column:party_b;not null" json:"party_b"`
	PhoneNumber            string    `gorm:"column:phone_number;not null" json:"phone_number"`
	AccountReference       string    `gorm:"column:account_reference" json:"account_reference"`
	TransactionDescription string    `gorm:"column:transaction_description" json:"transaction_description"`
	MerchantRequestID      string    `gorm:"column:merchant_request_id" json:"merchant_request_id"`
	CheckoutRequestID      string    `gorm:"column:checkout_request_id" json:"checkout_request_id"`
	ResponseCode           string    `gorm:"column:response_code" json:"response_code"`
	ResponseDescription    string    `gorm:"column:response_description" json:"response_description"`
	CustomerMessage        string    `gorm:"column:customer_message" json:"customer_message"`
	ResultCode             string    `gorm:"column:result_code" json:"result_code"`
	ResultDescription      string    `gorm:"column:result_description" json:"result_description"`
	MpesaReceiptNumber     string    `gorm:"column:mpesa_receipt_number" json:"mpesa_receipt_number"`
	TransactionDate        string    `gorm:"column:transaction_date" json:"transaction_date"`
	CreatedAt              time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName MpesaExpress's table name
func (*MpesaExpress) TableName() string {
	return TableNameMpesaExpress
}
