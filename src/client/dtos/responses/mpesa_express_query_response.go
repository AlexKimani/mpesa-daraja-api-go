package responses

type MpesaExpressQueryResponse struct {
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResultCode          string `json:"ResultCode"`
	ResultDesc          string `json:"ResultDesc"`
}

func NewMpesaExpressQueryResponse(response MpesaExpressQueryResponse) MpesaExpressQueryResponse {
	return MpesaExpressQueryResponse{
		ResponseCode:        response.ResponseCode,
		ResponseDescription: response.ResponseDescription,
		MerchantRequestID:   response.MerchantRequestID,
		CheckoutRequestID:   response.CheckoutRequestID,
		ResultCode:          response.ResultCode,
		ResultDesc:          response.ResultDesc,
	}
}
