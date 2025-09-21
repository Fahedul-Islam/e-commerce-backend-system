package users

import (
	"encoding/json"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/domain"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData domain.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid login data", http.StatusBadRequest)
		return
	}
	if err := h.srvc.UserValidate(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.srvc.AuthenticateUser(&loginData)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusUnauthorized)
		return
	}
	accessTokenString, refreshTokenString, err := h.generateToken(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]string{"access_token": accessTokenString, "refresh_token": refreshTokenString, "expires_in": h.cnf.JWT.TokenExpiry.String(), "token_type": "bearer"}, http.StatusOK)
}
