package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/paprykdev/homeledger/internal/auth"
	"github.com/paprykdev/homeledger/internal/helpers"
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
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var transaction models.Transaction

	if err := helpers.DecodeJSONRequest(r, &transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set user_id from authenticated user
	transaction.UserID = userID

	// Normalization & Validation
	if err := h.normalizeAndValidate(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now()

	transaction.ID = uuid.New().String()
	transaction.CreatedAt = now
	transaction.UpdatedAt = now

	// Sending query
	query := `
	INSERT INTO transactions(
		id,
		user_id,
		type,
		amount,
		description,
		category,
		account_id,
		created_at,
		updated_at
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err = h.DB.Exec(
		query,
		transaction.ID,
		transaction.UserID,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Category,
		transaction.AccountID,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)

	if err != nil {
		helpers.HandleDBError(w, "failed to create transaction", err)
		return
	}

	helpers.WriteJSONResponse(w, http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransactions(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var transactions []models.Transaction

	query := `
		SELECT id, user_id, type, amount, description, category, account_id, created_at, updated_at FROM transactions
		WHERE deleted_at IS NULL AND user_id = ?
	`

	rows, err := h.DB.Query(query, userID)

	if err != nil {
		helpers.HandleDBError(w, "failed to list transactions", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction

		if err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Category,
			&transaction.AccountID,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		); err != nil {
			helpers.HandleDBError(w, "failed to scan transaction", err)
			return
		}

		transactions = append(transactions, transaction)
	}

	// Check for errors from row iteration
	if err := rows.Err(); err != nil {
		helpers.HandleDBError(w, "error iterating transactions", err)
		return
	}

	helpers.WriteJSONResponse(w, http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	transaction, ok := h.getTransactionByIdOrError(w, id, userID)
	if !ok {
		return
	}

	helpers.WriteJSONResponse(w, http.StatusOK, transaction)
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	transaction, ok := h.getTransactionByIdOrError(w, id, userID)
	if !ok {
		return
	}

	deletedAt := time.Now()

	deleteQuery := `
		UPDATE transactions
		SET deleted_at = ?
		WHERE id = ?
	`

	if _, err := h.DB.Exec(deleteQuery, deletedAt, transaction.ID); err != nil {
		helpers.HandleDBError(w, "failed to delete transaction", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	transaction, ok := h.getTransactionByIdOrError(w, id, userID)
	if !ok {
		return
	}

	var updates models.Transaction
	if err := helpers.DecodeJSONRequest(r, &updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields if provided
	if updates.Type != "" {
		transaction.Type = updates.Type
	}
	if updates.Amount != 0 {
		transaction.Amount = updates.Amount
	}
	if updates.Description != "" {
		transaction.Description = updates.Description
	}
	if updates.Category != "" {
		transaction.Category = updates.Category
	}

	// Normalize & Validate
	if err := h.normalizeAndValidate(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction.UpdatedAt = time.Now()

	updateQuery := `
		UPDATE transactions
		SET type = ?, amount = ?, description = ?, category = ?, updated_at = ?
		WHERE id = ?
	`

	if _, err := h.DB.Exec(
		updateQuery,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Category,
		transaction.UpdatedAt,
		transaction.ID,
	); err != nil {
		helpers.HandleDBError(w, "failed to update transaction", err)
		return
	}

	helpers.WriteJSONResponse(w, http.StatusOK, transaction)
}

// Private funcs
func (h *TransactionHandler) getTransactionById(id string, userID string) (*models.Transaction, error) {
	var transaction models.Transaction

	query := `
	SELECT id, user_id, type, amount, description, category, account_id, created_at, updated_at FROM transactions
	WHERE id = ? AND user_id = ? AND deleted_at IS NULL
	`

	row := h.DB.QueryRow(query, id, userID)

	err := row.Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Description,
		&transaction.Category,
		&transaction.AccountID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (h *TransactionHandler) validateTransaction(transaction *models.Transaction) error {
	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	// Validate transaction type
	switch transaction.Type {
	case models.TransactionTypeExpense,
		models.TransactionTypeIncome:
	default:
		return errors.New("invalid type (income, expense)")
	}

	// Validate string field lengths
	if err := helpers.ValidateStringLength("category", transaction.Category, helpers.MaxStringFieldLength); err != nil {
		return err
	}

	if err := helpers.ValidateStringLength("description", transaction.Description, helpers.MaxDescriptionLength); err != nil {
		return err
	}

	if err := helpers.ValidateStringLength("account_id", transaction.AccountID, helpers.MaxStringFieldLength); err != nil {
		return err
	}

	// Validate that account_id is provided
	if transaction.AccountID == "" {
		return errors.New("account_id is required")
	}

	// Validate that account belongs to the user
	if err := h.validateAccountOwnership(transaction.UserID, transaction.AccountID); err != nil {
		return err
	}

	return nil
}

// validateAccountOwnership ensures the account belongs to the user
func (h *TransactionHandler) validateAccountOwnership(userID string, accountID string) error {
	query := `SELECT user_id FROM accounts WHERE id = ? AND deleted_at IS NULL`
	var ownerID string
	err := h.DB.QueryRow(query, accountID).Scan(&ownerID)

	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("account not found")
	}

	if err != nil {
		return errors.New("failed to validate account ownership")
	}

	if ownerID != userID {
		return errors.New("account does not belong to user")
	}

	return nil
}

func (h *TransactionHandler) normalizeTransaction(transaction *models.Transaction) {
	transaction.Category = strings.TrimSpace(transaction.Category)
	transaction.Category = strings.ToLower(transaction.Category)

	transaction.Type = models.TransactionType(strings.TrimSpace(string(transaction.Type)))
	transaction.Type = models.TransactionType(strings.ToLower(string(transaction.Type)))

	if transaction.Category == "" {
		transaction.Category = "other"
	}
}

func (h *TransactionHandler) getTransactionByIdOrError(w http.ResponseWriter, id string, userID string) (*models.Transaction, bool) {
	transaction, err := h.getTransactionById(id, userID)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	}
	if err != nil {
		helpers.HandleDBError(w, "failed to fetch transaction", err)
		return nil, false
	}
	return transaction, true
}

func (h *TransactionHandler) normalizeAndValidate(transaction *models.Transaction) error {
	h.normalizeTransaction(transaction)
	return h.validateTransaction(transaction)
}
