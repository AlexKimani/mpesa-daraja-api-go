package requests

type AccountBalanceRequest struct {
	OriginatorConversationID string `json:"-"`
	Initiator                string `json:"Initiator"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"Command ID"`
	PartyA                   string `json:"PartyA"`
	IdentifierType           string `json:"IdentifierType"`
	Remarks                  string `json:"Remarks"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	ResultURL                string `json:"ResultURL"`
}

func NewAccountBalanceRequest(request AccountBalanceRequest) AccountBalanceRequest {
	return AccountBalanceRequest{
		OriginatorConversationID: request.OriginatorConversationID,
		Initiator:                request.Initiator,
		SecurityCredential:       request.SecurityCredential,
		CommandID:                request.CommandID,
		PartyA:                   request.PartyA,
		IdentifierType:           request.IdentifierType,
		Remarks:                  request.Remarks,
		QueueTimeOutURL:          request.QueueTimeOutURL,
		ResultURL:                request.ResultURL,
	}
}
