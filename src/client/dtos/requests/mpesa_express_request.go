package requests

type MpesaExpressRequest struct {
	OriginatorConversationID string `json:"-"`
	BusinessShortCode        string `json:"BusinessShortCode"`
	Password                 string `json:"Password"`
	Timestamp                string `json:"Timestamp"`
	TransactionType          string `json:"TransactionType"`
	Amount                   string `json:"Amount"`
	PartyA                   string `json:"PartyA"`
	PartyB                   string `json:"PartyB"`
	PhoneNumber              string `json:"PhoneNumber"`
	CallBackURL              string `json:"CallBackURL"`
	AccountReference         string `json:"AccountReference"`
	TransactionDesc          string `json:"TransactionDesc"`
}