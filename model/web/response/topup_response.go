package response

import "time"

type TopupResponse struct {
	ID                    int       `json:"id"`
	UserID                int       `json:"user_id"`
	MidtransTransactionID string    `json:"midtrans_transaction_id"`
	MidtransOrderID       string    `json:"midtrans_order_id"`
	Amount                float64   `json:"amount"`
	Status                string    `json:"status"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
