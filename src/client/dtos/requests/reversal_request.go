package requests

type ReversalRequest struct {
	OriginatorConversationID string `json:"-"`
	Initiator                string `json:"Initiator"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"CommandID"`
	TransactionID            string `json:"TransactionID"`
	Amount                   string `json:"Amount"`
	ReceiverParty            string `json:"ReceiverParty"`
	RecieverIdentifierType   string `json:"RecieverIdentifierType"`
	ResultURL                string `json:"ResultURL"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	Remarks                  string `json:"Remarks"`
	Occasion                 string `json:"Occasion"`
}

func NewReversalRequest(request ReversalRequest) ReversalRequest {
	return ReversalRequest{
		OriginatorConversationID: request.OriginatorConversationID,
		Initiator:                request.Initiator,
		SecurityCredential:       request.SecurityCredential,
		CommandID:                request.CommandID,
		TransactionID:            request.TransactionID,
		Amount:                   request.Amount,
		ReceiverParty:            request.ReceiverParty,
		RecieverIdentifierType:   request.RecieverIdentifierType,
		ResultURL:                request.ResultURL,
		QueueTimeOutURL:          request.QueueTimeOutURL,
		Remarks:                  request.Remarks,
		Occasion:                 request.Occasion,
	}
}
