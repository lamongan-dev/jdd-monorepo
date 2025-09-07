package domain

type PaymentBody struct {
	Product     string  `json:"product"`
	Qty         int     `json:"qty"`
	Price       int     `json:"price"`
	Amount      float64 `json:"amount"`
	ReferenceId string  `json:"referenceId"`
	BuyerName   string  `json:"buyerName"`
	BuyerEmail  string  `json:"buyerEmail"`
}

type PaymentResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    PaymentData `json:"data"`
}

type PaymentData struct {
	SessionID string `json:"SessionID"`
	URL       string `json:"Url"`
}

type PaymentHeader struct {
	Va        string
	Signature string
	Timestamp string
	Url       string
}
