package requests

type C2BRegisterUrlRequest struct {
	RequestId       string `json:"-"`
	ShortCode       string `json:"ShortCode"`
	ResponseType    string `json:"ResponseType"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ValidationURL   string `json:"ValidationURL"`
}

func NewC2BRegisterUrlRequest(request C2BRegisterUrlRequest) C2BRegisterUrlRequest {
	return C2BRegisterUrlRequest{
		ShortCode:       request.ShortCode,
		ResponseType:    request.ResponseType,
		ConfirmationURL: request.ConfirmationURL,
		ValidationURL:   request.ValidationURL,
	}
}
