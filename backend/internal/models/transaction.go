package models

type Transaction struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
}
