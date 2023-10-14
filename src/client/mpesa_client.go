package client

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"mpesa-daraja-api-go/src/client/dtos/requests"
	"mpesa-daraja-api-go/src/client/dtos/responses"
	configs "mpesa-daraja-api-go/src/config"
	"net/http"
)

func DarajaAuth(config configs.Config) string {
	client := configs.NewRetryableClient(config)
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.Auth
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("Error preparing request service %+v", err)
		return ""
	}
	req.SetBasicAuth(config.Mpesa.ConsumerKey, config.Mpesa.ConsumerSecret)
	req.Header.Add("Accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Errorf("Error getting response from mpesa auth %+v", err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("Error getting response from mpesa auth %+v", err)
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Error reading response from mpesa auth %+v", err)
	}
	var authResponse responses.AuthorizationResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response from mpesa auth %+v", err)
		return ""
	}
	return authResponse.AccessToken
}

func MpesaApiCall(token string, url string, body bytes.Buffer, client *http.Client) []byte {
	req, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		log.Errorf("Error preparing mpesa request %+v", err)
		return nil
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Errorf("Error getting response from mpesa daraja %+v", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("Error getting response from mpesa daraja %+v", err)
		}
	}(response.Body)
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Error reading response from mpesa daraja %+v", err)
	}
	return responseBody
}

func MpesaAccountBalanceRequest(config configs.Config, request requests.AccountBalanceRequest) (responses.SuccessfulMpesaResponse, responses.FailedMpesaResponse) {
	successResponse := responses.NewSuccessfulMpesaResponse(responses.SuccessfulMpesaResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.Reversal
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling Account balance api request body %+v/n Error %+v", request, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.001.04",
			ErrorMessage: "Failed to marshal Account Balance Request",
		}
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa account balance api %+v", responseBody, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.002.04",
			ErrorMessage: "Failed to unmarshal account balance Response",
		}
	}
	return successResponse, failedResponse
}

func MpesaExpressRequest(config configs.Config, request requests.MpesaExpressRequest) responses.MpesaExpressResponse {
	checkoutResponse := responses.NewMpesaExpressResponse(responses.MpesaExpressResponse{})
	token := DarajaAuth(config)
	if token == "" {
		return checkoutResponse
	}
	client := configs.NewRetryableClient(config)
	var authToken = "Bearer " + token
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.Express
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling checkout request body %+v", err)
	}
	responseBody := MpesaApiCall(authToken, url, body, client)
	err = json.Unmarshal(responseBody, &checkoutResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response from mpesa checkout %+v", err)
	}
	return checkoutResponse
}

func MpesaExpressQueryRequest(config configs.Config, request requests.MpesaExpressQueryRequest) responses.MpesaExpressQueryResponse {
	checkoutQueryResponse := responses.NewMpesaExpressQueryResponse(responses.MpesaExpressQueryResponse{})
	token := DarajaAuth(config)
	if token == "" {
		return checkoutQueryResponse
	}
	client := configs.NewRetryableClient(config)
	var authToken = "Bearer " + token
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.Express
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling checkout query request body %+v", err)
	}
	responseBody := MpesaApiCall(authToken, url, body, client)
	err = json.Unmarshal(responseBody, &checkoutQueryResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response from mpesa checkout query %+v", err)
	}
	return checkoutQueryResponse
}

func MpesaC2BRegisterUrlRequest(config configs.Config, request requests.C2BRegisterUrlRequest) (responses.C2BRegisterUrlResponse, responses.FailedMpesaResponse) {
	c2bRegisterUrlResponse := responses.NewC2BRegisterUrlResponse(responses.C2BRegisterUrlResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return c2bRegisterUrlResponse, responses.FailedMpesaResponse{
			RequestId:    request.RequestId,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.C2bRegisterUrl
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling c2b register url request body %+v /nError %+v", request, err)
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &c2bRegisterUrlResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa c2b register url %+v", responseBody, err)
	}
	return c2bRegisterUrlResponse, failedResponse
}

func MpesaB2cRequest(config configs.Config, request requests.B2cRequest) (responses.SuccessfulMpesaResponse, responses.FailedMpesaResponse) {
	successResponse := responses.NewSuccessfulMpesaResponse(responses.SuccessfulMpesaResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.B2c
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling b2c api request body %+v/n Error %+v", request, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.001.04",
			ErrorMessage: "Failed to marshal B2c Request",
		}
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa b2c api %+v", responseBody, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.002.04",
			ErrorMessage: "Failed to unmarshal B2C Response",
		}
	}
	return successResponse, failedResponse
}

func MpesaTransactionStatusRequest(config configs.Config, request requests.TransactionStatusRequest) (responses.SuccessfulMpesaResponse, responses.FailedMpesaResponse) {
	successResponse := responses.NewSuccessfulMpesaResponse(responses.SuccessfulMpesaResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.TransactionStatus
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling transaction status api request body %+v/n Error %+v", request, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.001.04",
			ErrorMessage: "Failed to marshal B2c Request",
		}
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa transaction status api %+v", responseBody, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.002.04",
			ErrorMessage: "Failed to unmarshal B2C Response",
		}
	}
	return successResponse, failedResponse
}

func MpesaBusinessBuyGoodsRequest(config configs.Config, request requests.BusinessBuyGoodsRequest) (responses.SuccessfulMpesaResponse, responses.FailedMpesaResponse) {
	successResponse := responses.NewSuccessfulMpesaResponse(responses.SuccessfulMpesaResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.BusinessBuyGoods
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling business buy goods api request body %+v/n Error %+v", request, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.001.04",
			ErrorMessage: "Failed to marshal business buy goods Request",
		}
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa business buy goods api %+v", responseBody, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.002.04",
			ErrorMessage: "Failed to unmarshal business buy goods Response",
		}
	}
	return successResponse, failedResponse
}

func MpesaBusinessPayBillRequest(config configs.Config, request requests.BusinessPayBillRequest) (responses.SuccessfulMpesaResponse, responses.FailedMpesaResponse) {
	successResponse := responses.NewSuccessfulMpesaResponse(responses.SuccessfulMpesaResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.BusinessPayBill
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling business pay bill api request body %+v/n Error %+v", request, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.001.04",
			ErrorMessage: "Failed to marshal business pay bill Request",
		}
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa business pay bill api %+v", responseBody, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.002.04",
			ErrorMessage: "Failed to unmarshal business pay bill Response",
		}
	}
	return successResponse, failedResponse
}

func MpesaReversalRequest(config configs.Config, request requests.ReversalRequest) (responses.SuccessfulMpesaResponse, responses.FailedMpesaResponse) {
	successResponse := responses.NewSuccessfulMpesaResponse(responses.SuccessfulMpesaResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.Reversal
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling Reversal api request body %+v/n Error %+v", request, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.001.04",
			ErrorMessage: "Failed to marshal Reversal Request",
		}
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa reversal api %+v", responseBody, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.002.04",
			ErrorMessage: "Failed to unmarshal reversal Response",
		}
	}
	return successResponse, failedResponse
}

func MpesaTaxRemittanceRequest(config configs.Config, request requests.TaxRemittanceRequest) (responses.SuccessfulMpesaResponse, responses.FailedMpesaResponse) {
	successResponse := responses.NewSuccessfulMpesaResponse(responses.SuccessfulMpesaResponse{})
	failedResponse := responses.NewFailedMpesaResponse(responses.FailedMpesaResponse{})
	authToken := DarajaAuth(config)
	if authToken == "" {
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "404.001.04",
			ErrorMessage: "Invalid Access Token",
		}
	}
	client := configs.NewRetryableClient(config)
	var token = "Bearer " + authToken
	var url = config.Mpesa.BaseUrl + config.Mpesa.Urls.Reversal
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		log.Errorf("Error Marshalling Tax Remittance api request body %+v/n Error %+v", request, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.001.04",
			ErrorMessage: "Failed to marshal Tax Remittance Request",
		}
	}

	responseBody := MpesaApiCall(token, url, body, client)
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		log.Errorf("Error unmarshalling response %+v/n from mpesa Tax Remittance api %+v", responseBody, err)
		return successResponse, responses.FailedMpesaResponse{
			RequestId:    request.OriginatorConversationID,
			ErrorCode:    "401.002.04",
			ErrorMessage: "Failed to unmarshal Tax Remittance Response",
		}
	}
	return successResponse, failedResponse
}
