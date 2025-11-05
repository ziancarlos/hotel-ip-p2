package domain

import "time"

type Topup struct {
	ID                    int       `json:"id" db:"id"`
	UserID                int       `json:"user_id" db:"user_id"`
	MidtransTransactionID string    `json:"midtrans_transaction_id" db:"midtrans_transaction_id"`
	MidtransOrderID       string    `json:"midtrans_order_id" db:"midtrans_order_id"`
	Amount                float64   `json:"amount" db:"amount"`
	Status                string    `json:"status" db:"status"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
}
