package authHandler

import (
	"encoding/json"
	"net/http"

	authmodels "github.com/Atharv7901/E2EEChatApp/models"
)

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req authmodels.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.authService.CreateUser(req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
