package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/paprykdev/homeledger/internal/auth"
	"github.com/paprykdev/homeledger/internal/helpers"
	"github.com/paprykdev/homeledger/internal/models"
)

type AccountHandler struct {
	DB *sql.DB
}

func NewAccountHandler(db *sql.DB) *AccountHandler {
	return &AccountHandler{
		DB: db,
	}
}

// CreateAccount creates a new account for the authenticated user
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var account models.Account

	if err := helpers.DecodeJSONRequest(r, &account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if err := h.validateAccount(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set user_id and create ID
	account.UserID = userID
	account.ID = uuid.New().String()
	now := time.Now()
	account.CreatedAt = now
	account.UpdatedAt = now

	query := `
	INSERT INTO accounts (id, user_id, name, currency, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = h.DB.Exec(query, account.ID, account.UserID, account.Name, account.Currency, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		helpers.HandleDBError(w, "failed to create account", err)
		return
	}

	helpers.WriteJSONResponse(w, http.StatusCreated, account)
}

// GetAccounts returns all accounts for the authenticated user
func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var accounts []models.Account

	query := `
		SELECT id, user_id, name, currency, created_at, updated_at FROM accounts
		WHERE deleted_at IS NULL AND user_id = ?
		ORDER BY created_at DESC
	`

	rows, err := h.DB.Query(query, userID)
	if err != nil {
		helpers.HandleDBError(w, "failed to list accounts", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var account models.Account

		if err := rows.Scan(
			&account.ID,
			&account.UserID,
			&account.Name,
			&account.Currency,
			&account.CreatedAt,
			&account.UpdatedAt,
		); err != nil {
			helpers.HandleDBError(w, "failed to scan account", err)
			return
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		helpers.HandleDBError(w, "error iterating accounts", err)
		return
	}

	helpers.WriteJSONResponse(w, http.StatusOK, accounts)
}

// GetAccountById returns a specific account if it belongs to the user
func (h *AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	account, ok := h.getAccountByIdOrError(w, id, userID)
	if !ok {
		return
	}

	helpers.WriteJSONResponse(w, http.StatusOK, account)
}

// DeleteAccount soft-deletes an account
func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from context
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	account, ok := h.getAccountByIdOrError(w, id, userID)
	if !ok {
		return
	}

	deletedAt := time.Now()

	deleteQuery := `
		UPDATE accounts
		SET deleted_at = ?
		WHERE id = ?
	`

	if _, err := h.DB.Exec(deleteQuery, deletedAt, account.ID); err != nil {
		helpers.HandleDBError(w, "failed to delete account", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Private helper methods

func (h *AccountHandler) getAccountByIdOrError(w http.ResponseWriter, id string, userID string) (*models.Account, bool) {
	account, err := h.getAccountById(id, userID)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	}
	if err != nil {
		helpers.HandleDBError(w, "failed to fetch account", err)
		return nil, false
	}
	return account, true
}

func (h *AccountHandler) getAccountById(id string, userID string) (*models.Account, error) {
	var account models.Account

	query := `
	SELECT id, user_id, name, currency, created_at, updated_at FROM accounts
	WHERE id = ? AND user_id = ? AND deleted_at IS NULL
	`

	row := h.DB.QueryRow(query, id, userID)

	err := row.Scan(
		&account.ID,
		&account.UserID,
		&account.Name,
		&account.Currency,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (h *AccountHandler) validateAccount(account *models.Account) error {
	if err := helpers.ValidateStringLength("name", account.Name, helpers.MaxStringFieldLength); err != nil {
		return err
	}

	if account.Name == "" {
		return errors.New("name is required")
	}

	// Validate currency
	validCurrencies := map[string]bool{
		"PLN": true,
		"USD": true,
		"EUR": true,
	}

	if !validCurrencies[account.Currency] {
		return errors.New("invalid currency (PLN, USD, EUR)")
	}

	return nil
}
