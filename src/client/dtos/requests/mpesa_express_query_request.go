package requests

type MpesaExpressQueryRequest struct {
	OriginatorConversationID string `json:"-"`
	BusinessShortCode        string `json:"BusinessShortCode"`
	Password                 string `json:"Password"`
	Timestamp                string `json:"Timestamp"`
	CheckoutRequestID        string `json:"CheckoutRequestID"`
}

func NewMpesaExpressQueryRequest(request MpesaExpressQueryRequest) MpesaExpressQueryRequest {
	return MpesaExpressQueryRequest{
		OriginatorConversationID: request.OriginatorConversationID,
		BusinessShortCode:        request.BusinessShortCode,
		Password:                 request.Password,
		Timestamp:                request.Timestamp,
		CheckoutRequestID:        request.CheckoutRequestID,
	}
}
