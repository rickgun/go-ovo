package ovo

type Header struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

type PushToPayRequest struct {
	Type					string `json:"type"`
	ProcessingCode			string `json:"processingCode"`
	Amount					float64 `json:"amount"`
	Date					string `json:"date"`
	ReferenceNumber			string `json:"referenceNumber"`
	TID						string `json:"tid"`
	MID						string `json:"mid"`
	MerchantID				string `json:"merchantId"`
	StoreCode				string `json:"storeCode"`
	AppSource				string `json:"appSource"`
	TransactionRequestData	TransactionRequestData `json:"transactionRequestData"`
}

type TransactionRequestData struct {
	BatchNo			string `json:"batchNo"`
	MerchantInvoice string `json:"merchantInvoice,omitempty"`
	Phone			string `json:"phone,omitempty"`
}