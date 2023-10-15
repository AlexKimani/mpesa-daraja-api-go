package model

import (
	"time"
)

const TableNameInitiator = "initiator"

// Initiator mapped from table <initiator>
type Initiator struct {
	ID                  int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	InitiatorName       string    `gorm:"column:initiator_name;not null" json:"initiator_name"`
	InitiatorCredential string    `gorm:"column:initiator_credential;not null" json:"initiator_credential"`
	CreatedAt           time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Initiator's table name
func (*Initiator) TableName() string {
	return TableNameInitiator
}

func NewInitiator(initiator Initiator) Initiator {
	return Initiator{
		ID:                  initiator.ID,
		InitiatorName:       initiator.InitiatorName,
		InitiatorCredential: initiator.InitiatorCredential,
		CreatedAt:           initiator.CreatedAt,
		UpdatedAt:           initiator.UpdatedAt,
	}
}
