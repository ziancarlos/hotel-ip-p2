package domain

import "time"

// User represents a user domain model (database entity)
type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`      // Password hash, not exposed in JSON
	Balance   float64   `json:"balance" db:"balance"` // Balance in Indonesian Rupiah (IDR)
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
