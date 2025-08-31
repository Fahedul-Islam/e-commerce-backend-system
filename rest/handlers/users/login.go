package users

import (
	"encoding/json"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData database.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid login data", http.StatusBadRequest)
		return
	}
	if err := (loginData.Validate()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.Repo.Authenticate(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusUnauthorized)
		return
	}
	tokenString, err := h.generateToken(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]string{"token": tokenString, "expires_in": h.Repo.TokenExpiry.String(), "token_type": "bearer"}, http.StatusOK)
}
