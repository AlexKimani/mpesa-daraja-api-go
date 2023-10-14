package results

type MpesaAccountBalanceResult struct {
	Result struct {
		ResultType               string `json:"ResultType"`
		ResultCode               string `json:"ResultCode"`
		ResultDesc               string `json:"ResultDesc"`
		OriginatorConversationID string `json:"OriginatorConversationID"`
		ConversationID           string `json:"ConversationID"`
		TransactionID            string `json:"TransactionID"`
		ResultParameter          struct {
			ResultParameters []struct {
				Key   string `json:"Key"`
				Value string `json:"Value"`
			} `json:"ResultParameters"`
		} `json:"ResultParameter"`
		ReferenceData struct {
			ReferenceItem struct {
				Key   string `json:"Key"`
				Value string `json:"Value"`
			} `json:"ReferenceItem"`
		} `json:"ReferenceData"`
	} `json:"Result"`
}
