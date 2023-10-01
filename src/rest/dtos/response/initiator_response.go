package response

import "time"

type InitiatorResponse struct {
	ID                  int64     `json:"id,omitempty"`
	InitiatorName       string    `json:"initiator_name,omitempty"`
	InitiatorCredential string    `json:"initiator_credential,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
