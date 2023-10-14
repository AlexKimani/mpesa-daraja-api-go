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
