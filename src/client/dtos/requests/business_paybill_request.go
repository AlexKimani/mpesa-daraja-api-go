package requests

type BusinessPayBillRequest struct {
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
	Requester                string `json:"Requester"`
	Remarks                  string `json:"Remarks"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	ResultURL                string `json:"ResultURL"`
}

func NewBusinessPayBillRequest(request BusinessPayBillRequest) BusinessPayBillRequest {
	return BusinessPayBillRequest{
		Initiator:              request.Initiator,
		SecurityCredential:     request.SecurityCredential,
		CommandID:              request.CommandID,
		SenderIdentifierType:   request.SenderIdentifierType,
		RecieverIdentifierType: request.RecieverIdentifierType,
		Amount:                 request.Amount,
		PartyA:                 request.PartyA,
		PartyB:                 request.PartyB,
		AccountReference:       request.AccountReference,
		Requester:              request.Requester,
		Remarks:                request.Remarks,
		QueueTimeOutURL:        request.QueueTimeOutURL,
		ResultURL:              request.ResultURL,
	}
}
