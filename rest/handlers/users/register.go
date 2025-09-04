package users

import (
	"encoding/json"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser database.UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := newUser.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := util.ValidatePassword(newUser.Password); err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	// Hash password before storing
	hashedPassword, _ := util.HashPassword(newUser.Password)

	// Generate OTP
	otp, _ := util.GenerateOTP()

	// Store user + OTP in Redis temporarily
	tempUser := map[string]string{
		"username": newUser.Username,
		"email":    newUser.Email,
		"password": hashedPassword,
		"roles":    newUser.Roles,
		"otp":      otp,
	}
	err := database.SaveTempUser(newUser.Email, tempUser) // save to Redis
	if err != nil {
		http.Error(w, "Failed to save OTP", http.StatusInternalServerError)
		return
	}

	// Send OTP email
	_ = util.SendOTPEmail(newUser.Email, otp)

	util.SendData(w, map[string]string{
		"message": "OTP sent to your email. Verify to complete registration.",
	}, http.StatusOK)
}
