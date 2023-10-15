package model

import (
	"time"
)

const TableNameCustomerToBusiness = "customer_to_business"

// CustomerToBusiness mapped from table <customer_to_business>
type CustomerToBusiness struct {
	ID                   int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Msisdn               string    `gorm:"column:msisdn;not null" json:"msisdn"`
	TransactionType      string    `gorm:"column:transaction_type;not null" json:"transaction_type"`
	TransactionID        string    `gorm:"column:transaction_id;not null" json:"transaction_id"`
	TransactionAmount    float64   `gorm:"column:transaction_amount" json:"transaction_amount"`
	MpesaTransactionID   string    `gorm:"column:mpesa_transaction_id;not null" json:"mpesa_transaction_id"`
	MpesaTransactionTime string    `gorm:"column:mpesa_transaction_time;not null" json:"mpesa_transaction_time"`
	BusinessShortCode    string    `gorm:"column:business_short_code;not null" json:"business_short_code"`
	BillReferenceNumber  string    `gorm:"column:bill_reference_number;not null" json:"bill_reference_number"`
	InvoiceNumber        string    `gorm:"column:invoice_number" json:"invoice_number"`
	OrgAccountBalance    float64   `gorm:"column:org_account_balance" json:"org_account_balance"`
	ThirdPartyTransID    string    `gorm:"column:third_party_trans_id" json:"third_party_trans_id"`
	FirstName            string    `gorm:"column:first_name" json:"first_name"`
	MiddleName           string    `gorm:"column:middle_name" json:"middle_name"`
	LastName             string    `gorm:"column:last_name" json:"last_name"`
	ResultCode           string    `gorm:"column:result_code" json:"result_code"`
	ResultDescription    string    `gorm:"column:result_description" json:"result_description"`
	CreatedAt            time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName CustomerToBusiness's table name
func (*CustomerToBusiness) TableName() string {
	return TableNameCustomerToBusiness
}
