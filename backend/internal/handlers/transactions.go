package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/paprykdev/homeledger/internal/models"
)

type TransactionHandler struct {
	DB *sql.DB
}

func NewTransactionHandler(db *sql.DB) *TransactionHandler {
	return &TransactionHandler{
		DB: db,
	}
}

func (h *TransactionHandler) CreateTransaction(
	w http.ResponseWriter,
	r *http.Request,
) {
	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if transaction.Amount <= 0 {
		http.Error(w, "amount have to be greater than 0", http.StatusBadRequest)
		return
	}

	transaction.ID = uuid.New().String()
	transaction.CreatedAt = time.Now().Format(time.RFC3339)

	query := `
	INSERT INTO transactions(
		id,
		amount,
		description,
		created_at
	)
	VALUES (?, ?, ?, ?)
	`

	_, err = h.DB.Exec(
		query,
		transaction.ID,
		transaction.Amount,
		transaction.Description,
		transaction.CreatedAt,
	)

	if err != nil {
		http.Error(
			w,
			"failed to create transaction",
			http.StatusInternalServerError,
		)
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) GetTransactions(
	w http.ResponseWriter,
	r *http.Request,
) {
	var transactions []models.Transaction

	query := `
		SELECT id, amount, description, created_at FROM transactions
		WHERE deleted_at IS NULL
	`

	rows, err := h.DB.Query(query)

	if err != nil {
		http.Error(
			w,
			"failed to list transactions",
			http.StatusInternalServerError,
		)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.Amount,
			&transaction.Description,
			&transaction.CreatedAt,
		)
		if err != nil {
			http.Error(w, "failed to scan", http.StatusInternalServerError)
			return
		}

		transactions = append(transactions, transaction)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	transaction, err := h.getTransactionById(id)

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to fetch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	transaction, err := h.getTransactionById(id)

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to fetch", http.StatusInternalServerError)
		return
	}

	deletedAt := time.Now().Format(time.RFC3339)

	deleteQuery := `
		UPDATE transactions
		SET deleted_at = ?
		WHERE id = ?
	`

	_, err = h.DB.Exec(deleteQuery, deletedAt, transaction.ID)
	if err != nil {
		http.Error(w, "error fetching", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


// Private funcs
func (h *TransactionHandler) getTransactionById(id string) (*models.Transaction, error) {
	var transaction models.Transaction

	query := `
	SELECT id, amount, created_at, description FROM transactions WHERE id = ? AND deleted_at IS NULL
	`

	row := h.DB.QueryRow(query, id)

	err := row.Scan(
		&transaction.ID,
		&transaction.Amount,
		&transaction.CreatedAt,
		&transaction.Description,
	)

	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
