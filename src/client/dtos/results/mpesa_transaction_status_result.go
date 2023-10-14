package results

type MpesaTransactionStatusResult struct {
	Result struct {
		ConversationID           string `json:"ConversationID"`
		OriginatorConversationID string `json:"OriginatorConversationID"`
		ReferenceData            struct {
			ReferenceItem struct {
				Key string `json:"Key"`
			} `json:"ReferenceItem"`
		} `json:"ReferenceData"`
		ResultCode       int    `json:"ResultCode"`
		ResultDesc       string `json:"ResultDesc"`
		ResultParameters struct {
			ResultParameter []struct {
				Key   string `json:"Key"`
				Value string `json:"Value,omitempty"`
			} `json:"ResultParameter"`
		} `json:"ResultParameters"`
		ResultType    int    `json:"ResultType"`
		TransactionID string `json:"TransactionID"`
	} `json:"Result"`
}
