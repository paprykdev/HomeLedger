package models

import "time"

type Account struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Currency  string     `json:"currency"`
	UserID    string     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
