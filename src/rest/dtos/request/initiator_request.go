package request

type InitiatorRequest struct {
	InitiatorName       string `json:"initiator_name"`
	InitiatorCredential string `json:"initiator_credential"`
}
