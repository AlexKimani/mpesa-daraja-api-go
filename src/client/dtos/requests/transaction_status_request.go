package requests

type TransactionStatusRequest struct {
	Initiator                string `json:"Initiator"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"Command ID"`
	TransactionID            string `json:"Transaction ID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	PartyA                   string `json:"PartyA"`
	IdentifierType           string `json:"IdentifierType"`
	ResultURL                string `json:"ResultURL"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	Remarks                  string `json:"Remarks"`
	Occasion                 string `json:"Occasion"`
}

func NewTransactionStatusRequest(request TransactionStatusRequest) TransactionStatusRequest {
	return TransactionStatusRequest{
		Initiator:                request.Initiator,
		SecurityCredential:       request.SecurityCredential,
		CommandID:                request.CommandID,
		TransactionID:            request.TransactionID,
		OriginatorConversationID: request.OriginatorConversationID,
		PartyA:                   request.PartyA,
		IdentifierType:           request.IdentifierType,
		ResultURL:                request.ResultURL,
		QueueTimeOutURL:          request.QueueTimeOutURL,
		Remarks:                  request.Remarks,
		Occasion:                 request.Occasion,
	}
}
