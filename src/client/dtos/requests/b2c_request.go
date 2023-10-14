package requests

type B2cRequest struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	InitiatorName            string `json:"InitiatorName"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"CommandID"`
	Amount                   string `json:"Amount"`
	PartyA                   string `json:"PartyA"`
	PartyB                   string `json:"PartyB"`
	Remarks                  string `json:"Remarks"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	ResultURL                string `json:"ResultURL"`
	Occassion                string `json:"Occassion"`
}

func NewB2CRequest(request B2cRequest) B2cRequest {
	return B2cRequest{
		OriginatorConversationID: request.OriginatorConversationID,
		InitiatorName:            request.InitiatorName,
		SecurityCredential:       request.SecurityCredential,
		CommandID:                request.CommandID,
		Amount:                   request.Amount,
		PartyA:                   request.PartyA,
		PartyB:                   request.PartyB,
		Remarks:                  request.Remarks,
		QueueTimeOutURL:          request.QueueTimeOutURL,
		ResultURL:                request.ResultURL,
		Occassion:                request.Occassion,
	}
}
