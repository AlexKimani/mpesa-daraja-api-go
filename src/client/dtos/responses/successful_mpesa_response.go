package responses

type SuccessfulMpesaResponse struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func NewSuccessfulMpesaResponse(response SuccessfulMpesaResponse) SuccessfulMpesaResponse {
	return SuccessfulMpesaResponse{
		ConversationID:           response.ConversationID,
		OriginatorConversationID: response.OriginatorConversationID,
		ResponseCode:             response.ResponseCode,
		ResponseDescription:      response.ResponseDescription,
	}
}
