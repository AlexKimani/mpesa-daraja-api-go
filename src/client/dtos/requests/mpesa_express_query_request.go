package requests

type MpesaExpressQueryRequest struct {
	OriginatorConversationID string `json:"-"`
	BusinessShortCode        string `json:"BusinessShortCode"`
	Password                 string `json:"Password"`
	Timestamp                string `json:"Timestamp"`
	CheckoutRequestID        string `json:"CheckoutRequestID"`
}
