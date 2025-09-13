package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Fahedul-Islam/e-commerce/database/connections"
	"github.com/Fahedul-Islam/e-commerce/database/repository"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *UserHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve user data from Redis
	tempUser, err := connections.GetTempUser(payload.Email)
	if err != nil {
		http.Error(w, "OTP expired or invalid", http.StatusUnauthorized)
		return
	}

	// Check OTP
	if tempUser["otp"] != payload.OTP {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	// Insert into DB
	user := repository.User{
		Username:     tempUser["username"],
		Email:        tempUser["email"],
		PasswordHash: tempUser["password"],
		Roles:        tempUser["roles"],
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := h.Repo.Create(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Remove Redis entry after successful verification
	_ = connections.DeleteTempUser(payload.Email)

	util.SendData(w, user, http.StatusCreated)
}
