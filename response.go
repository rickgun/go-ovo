package ovo

type PushToPayResponse struct {
	Type					string  `json:"type"`
	ProcessingCode			string  `json:"processingCode"`
	Amount					float64 `json:"amount"`
	Date					string  `json:"date"`
	TraceNumber				int     `json:"traceNumber"`
	HostTime				string  `json:"hostTime"`
	HostDate				string  `json:"hostDate"`
	ReferenceNumber			int     `json:"referenceNumber"`
	ApprovalCode			string  `json:"approvalCode,omitempty"`
	ResponseCode			string  `json:"responseCode"`
	TID						string  `json:"tid"`
	MID						string  `json:"mid"`
	Message					string  `json:"message,omitempty"`
	TransactionRequestData	TransactionRequestData `json:"transactionRequestData"`
	TransactionResponseData	TransactionResponseData `json:"transactionResponseData,omitempty"`
}

type TransactionResponseData struct {
	StoreAddress1    string `json:"storeAddress1"`
	StoreAddress2    string `json:"storeAddress2"`
	PaymentType      string `json:"paymentType"`
	OvoID            string `json:"ovoid"`
	FullName         string `json:"fullName"`
	StoreCode        string `json:"storeCode"`
	StoreName		 string `json:"storeName"`
	CashUsed         string `json:"cashUsed"`
	CashBalance      string `json:"cashBalance,omitempty"`
	OvoPointsEarned  string `json:"ovoPointsEarned,omitempty"`
	OvoPointsUsed	 string `json:"ovoPointsUsed"`
	OvoPointsBalance string `json:"ovoPointsBalance,omitempty"`
}