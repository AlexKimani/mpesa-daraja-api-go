package responses

type MpesaExpressResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}

func NewMpesaExpressResponse(response MpesaExpressResponse) MpesaExpressResponse {
	return MpesaExpressResponse{
		MerchantRequestID:   response.MerchantRequestID,
		CheckoutRequestID:   response.CheckoutRequestID,
		ResponseCode:        response.ResponseCode,
		ResponseDescription: response.ResponseDescription,
		CustomerMessage:     response.CustomerMessage,
	}
}
