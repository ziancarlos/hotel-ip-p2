package request

type TopupWebhookRequest struct {
	TransactionStatus string `json:"transaction_status"`
	StatusCode        string `json:"status_code"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	SignatureKey      string `json:"signature_key"`
	UserID            int    `json:"user_id"`
}
