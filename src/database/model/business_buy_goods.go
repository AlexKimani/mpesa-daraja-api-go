package model

import (
	"time"
)

const TableNameBusinessBuyGood = "business_buy_goods"

// BusinessBuyGood mapped from table <business_buy_goods>
type BusinessBuyGood struct {
	ID                               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	InitiatorID                      int64     `gorm:"column:initiator_id" json:"initiator_id"`
	CommandID                        string    `gorm:"column:command_id;not null" json:"command_id"`
	SenderIdentifierType             string    `gorm:"column:sender_identifier_type" json:"sender_identifier_type"`
	ReceiverIdentifierType           string    `gorm:"column:receiver_identifier_type" json:"receiver_identifier_type"`
	TransactionAmount                float64   `gorm:"column:transaction_amount" json:"transaction_amount"`
	PartyA                           string    `gorm:"column:party_a;not null" json:"party_a"`
	PartyB                           string    `gorm:"column:party_b;not null" json:"party_b"`
	Remarks                          string    `gorm:"column:remarks" json:"remarks"`
	AccountReference                 string    `gorm:"column:account_reference" json:"account_reference"`
	Requester                        string    `gorm:"column:requester" json:"requester"`
	OriginatorConversationID         string    `gorm:"column:originator_conversation_id;not null" json:"originator_conversation_id"`
	ConversationID                   string    `gorm:"column:conversation_id;not null" json:"conversation_id"`
	ResponseCode                     string    `gorm:"column:response_code" json:"response_code"`
	ResponseDescription              string    `gorm:"column:response_description" json:"response_description"`
	ResultType                       string    `gorm:"column:result_type" json:"result_type"`
	ResultCode                       string    `gorm:"column:result_code" json:"result_code"`
	ResultDescription                string    `gorm:"column:result_description" json:"result_description"`
	MpesaTransactionID               string    `gorm:"column:mpesa_transaction_id" json:"mpesa_transaction_id"`
	DebitAccountBalance              string    `gorm:"column:debit_account_balance" json:"debit_account_balance"`
	Amount                           float64   `gorm:"column:amount" json:"amount"`
	DebitPartyAffectedAccountBalance string    `gorm:"column:debit_party_affected_account_balance" json:"debit_party_affected_account_balance"`
	MpesaTransactionCompletedTime    string    `gorm:"column:mpesa_transaction_completed_time" json:"mpesa_transaction_completed_time"`
	DebitPartyCharges                float64   `gorm:"column:debit_party_charges" json:"debit_party_charges"`
	ReceiverPartyPublicName          string    `gorm:"column:receiver_party_public_name" json:"receiver_party_public_name"`
	Currency                         string    `gorm:"column:currency" json:"currency"`
	InitiatorAccountCurrentBalance   string    `gorm:"column:initiator_account_current_balance" json:"initiator_account_current_balance"`
	BillReferenceNumber              string    `gorm:"column:bill_reference_number" json:"bill_reference_number"`
	ErrorCompletedTime               string    `gorm:"column:error_completed_time" json:"error_completed_time"`
	CreatedAt                        time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt                        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName BusinessBuyGood's table name
func (*BusinessBuyGood) TableName() string {
	return TableNameBusinessBuyGood
}
