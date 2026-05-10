package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

const (
	// MaxRequestSize limits JSON request body to 1MB
	MaxRequestSize = 1 << 20 // 1MB
	// MaxStringFieldLength limits string fields to 1000 characters
	MaxStringFieldLength = 1000
	// MaxDescriptionLength limits description to 5000 characters
	MaxDescriptionLength = 5000
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	if err := writeJSON(w, status, data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		log.Println("error writing JSON response:", err)
	}
}

func DecodeJSONRequest(r *http.Request, data interface{}) error {
	// Limit request body size to prevent DoS
	r.Body = http.MaxBytesReader(nil, r.Body, MaxRequestSize)

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		if err == io.EOF {
			return errors.New("empty request body")
		}
		if err.Error() == "http: request body too large" {
			return errors.New("request body too large")
		}
		log.Println("error decoding request:", err)
		return errors.New("invalid request body")
	}
	return nil
}

func HandleDBError(w http.ResponseWriter, message string, err error) {
	http.Error(w, message, http.StatusInternalServerError)
	log.Println(message+":", err)
}

// ValidateStringLength validates that a string field is within acceptable bounds
func ValidateStringLength(field string, value string, maxLength int) error {
	if len(value) > maxLength {
		return errors.New(field + " exceeds maximum length of " + string(rune(maxLength)))
	}
	return nil
}
