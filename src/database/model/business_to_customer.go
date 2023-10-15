package model

import (
	"time"
)

const TableNameBusinessToCustomer = "business_to_customer"

// BusinessToCustomer mapped from table <business_to_customer>
type BusinessToCustomer struct {
	ID                               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	InitiatorID                      int64     `gorm:"column:initiator_id" json:"initiator_id"`
	OriginatorConversationID         string    `gorm:"column:originator_conversation_id;not null" json:"originator_conversation_id"`
	CommandID                        string    `gorm:"column:command_id;not null" json:"command_id"`
	Amount                           float64   `gorm:"column:amount" json:"amount"`
	PartyA                           string    `gorm:"column:party_a;not null" json:"party_a"`
	PartyB                           string    `gorm:"column:party_b;not null" json:"party_b"`
	Remarks                          string    `gorm:"column:remarks" json:"remarks"`
	Occasion                         string    `gorm:"column:occasion" json:"occasion"`
	ConversationID                   string    `gorm:"column:conversation_id;not null" json:"conversation_id"`
	ResponseCode                     string    `gorm:"column:response_code" json:"response_code"`
	ResponseDescription              string    `gorm:"column:response_description" json:"response_description"`
	ErrorRequestID                   string    `gorm:"column:error_request_id" json:"error_request_id"`
	ErrorCode                        string    `gorm:"column:error_code" json:"error_code"`
	ErrorMessage                     string    `gorm:"column:error_message" json:"error_message"`
	ResultType                       string    `gorm:"column:result_type" json:"result_type"`
	ResultCode                       string    `gorm:"column:result_code" json:"result_code"`
	ResultDescription                string    `gorm:"column:result_description" json:"result_description"`
	MpesaTransactionID               string    `gorm:"column:mpesa_transaction_id" json:"mpesa_transaction_id"`
	TransactionAmount                float64   `gorm:"column:transaction_amount" json:"transaction_amount"`
	MpesaTransactionReceipt          string    `gorm:"column:mpesa_transaction_receipt" json:"mpesa_transaction_receipt"`
	IsRecipientRegisteredCustomer    string    `gorm:"column:is_recipient_registered_customer" json:"is_recipient_registered_customer"`
	ChargesPaidAccountAvailableFunds float64   `gorm:"column:charges_paid_account_available_funds" json:"charges_paid_account_available_funds"`
	ReceiverPartyPublicName          string    `gorm:"column:receiver_party_public_name" json:"receiver_party_public_name"`
	MpesaTransactionCompletedDate    time.Time `gorm:"column:mpesa_transaction_completed_date" json:"mpesa_transaction_completed_date"`
	UtilityAccountAvailableFunds     float64   `gorm:"column:utility_account_available_funds" json:"utility_account_available_funds"`
	WorkingAccountAvailableFunds     float64   `gorm:"column:working_account_available_funds" json:"working_account_available_funds"`
	CreatedAt                        time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt                        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName BusinessToCustomer's table name
func (*BusinessToCustomer) TableName() string {
	return TableNameBusinessToCustomer
}
