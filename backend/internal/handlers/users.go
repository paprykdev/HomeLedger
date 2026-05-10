package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/paprykdev/homeledger/internal/auth"
	"github.com/paprykdev/homeledger/internal/helpers"
	"github.com/paprykdev/homeledger/internal/models"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string        `json:"token"`
	User  models.User   `json:"user"`
}

// Register creates a new user
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := helpers.DecodeJSONRequest(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if err := h.validateRegisterRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	passwordHash, err := auth.HashPassword(req.Password)
	if err != nil {
		helpers.HandleDBError(w, "failed to hash password", err)
		return
	}

	// Create user
	user := &models.User{
		ID:        uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `
	INSERT INTO users (id, username, email, password_hash, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = h.DB.Exec(query, user.ID, user.Username, user.Email, passwordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		// Check for duplicate username/email
		if err.Error() == "UNIQUE constraint failed: users.username" {
			http.Error(w, "username already exists", http.StatusConflict)
		} else if err.Error() == "UNIQUE constraint failed: users.email" {
			http.Error(w, "email already exists", http.StatusConflict)
		} else {
			helpers.HandleDBError(w, "failed to create user", err)
		}
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, 24)
	if err != nil {
		helpers.HandleDBError(w, "failed to generate token", err)
		return
	}

	response := AuthResponse{
		Token: token,
		User:  *user,
	}

	helpers.WriteJSONResponse(w, http.StatusCreated, response)
}

// Login authenticates a user and returns a token
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := helpers.DecodeJSONRequest(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" || req.Password == "" {
		http.Error(w, "username and password required", http.StatusBadRequest)
		return
	}

	// Find user by username
	var user models.User
	var passwordHash string

	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE username = ? AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, req.Username).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if err != nil {
		helpers.HandleDBError(w, "failed to fetch user", err)
		return
	}

	// Get password hash for verification
	hashQuery := `SELECT password_hash FROM users WHERE id = ?`
	err = h.DB.QueryRow(hashQuery, user.ID).Scan(&passwordHash)
	if err != nil {
		helpers.HandleDBError(w, "failed to fetch password hash", err)
		return
	}

	// Verify password
	if !auth.VerifyPassword(passwordHash, req.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, 24)
	if err != nil {
		helpers.HandleDBError(w, "failed to generate token", err)
		return
	}

	response := AuthResponse{
		Token: token,
		User:  user,
	}

	helpers.WriteJSONResponse(w, http.StatusOK, response)
}

// validateRegisterRequest validates registration input
func (h *UserHandler) validateRegisterRequest(req *RegisterRequest) error {
	if err := helpers.ValidateStringLength("username", req.Username, helpers.MaxStringFieldLength); err != nil {
		return err
	}

	if err := helpers.ValidateStringLength("email", req.Email, helpers.MaxStringFieldLength); err != nil {
		return err
	}

	if err := helpers.ValidateStringLength("password", req.Password, helpers.MaxStringFieldLength); err != nil {
		return err
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if req.Username == "" {
		return errors.New("username is required")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
