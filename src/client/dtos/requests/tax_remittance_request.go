package requests

type TaxRemittanceRequest struct {
	OriginatorConversationID string `json:"-"`
	Initiator                string `json:"Initiator"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"Command ID"`
	SenderIdentifierType     string `json:"SenderIdentifierType"`
	RecieverIdentifierType   string `json:"RecieverIdentifierType"`
	Amount                   string `json:"Amount"`
	PartyA                   string `json:"PartyA"`
	PartyB                   string `json:"PartyB"`
	AccountReference         string `json:"AccountReference"`
	Remarks                  string `json:"Remarks"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	ResultURL                string `json:"ResultURL"`
}

func NewTaxRemittanceRequest(request TaxRemittanceRequest) TaxRemittanceRequest {
	return TaxRemittanceRequest{
		OriginatorConversationID: request.OriginatorConversationID,
		Initiator:                request.Initiator,
		SecurityCredential:       request.SecurityCredential,
		CommandID:                request.CommandID,
		SenderIdentifierType:     request.SenderIdentifierType,
		RecieverIdentifierType:   request.RecieverIdentifierType,
		Amount:                   request.Amount,
		PartyA:                   request.PartyA,
		PartyB:                   request.PartyB,
		AccountReference:         request.AccountReference,
		Remarks:                  request.Remarks,
		QueueTimeOutURL:          request.QueueTimeOutURL,
		ResultURL:                request.ResultURL,
	}
}
