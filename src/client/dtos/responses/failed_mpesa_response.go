package responses

type FailedMpesaResponse struct {
	RequestId    string `json:"requestId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewFailedMpesaResponse(response FailedMpesaResponse) FailedMpesaResponse {
	return FailedMpesaResponse{
		RequestId:    response.RequestId,
		ErrorCode:    response.ErrorCode,
		ErrorMessage: response.ErrorMessage,
	}
}
