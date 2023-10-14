package responses

type C2BRegisterUrlResponse struct {
	OriginatorCoversationID string `json:"OriginatorCoversationID"`
	ResponseCode            string `json:"ResponseCode"`
	ResponseDescription     string `json:"ResponseDescription"`
}

func NewC2BRegisterUrlResponse(response C2BRegisterUrlResponse) C2BRegisterUrlResponse {
	return C2BRegisterUrlResponse{
		OriginatorCoversationID: response.OriginatorCoversationID,
		ResponseCode:            response.ResponseCode,
		ResponseDescription:     response.ResponseCode,
	}
}
