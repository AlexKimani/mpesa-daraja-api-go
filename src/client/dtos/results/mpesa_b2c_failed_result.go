package results

type MpesaB2cFailedResult struct {
	Result struct {
		ResultType               int    `json:"ResultType"`
		ResultCode               int    `json:"ResultCode"`
		ResultDesc               string `json:"ResultDesc"`
		OriginatorConversationID string `json:"OriginatorConversationID"`
		ConversationID           string `json:"ConversationID"`
		TransactionID            string `json:"TransactionID"`
		ReferenceData            struct {
			ReferenceItem struct {
				Key   string `json:"Key"`
				Value string `json:"Value"`
			} `json:"ReferenceItem"`
		} `json:"ReferenceData"`
	} `json:"Result"`
}
