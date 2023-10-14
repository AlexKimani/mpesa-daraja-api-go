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
