package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser database.UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}
	if err := (newUser.Validate()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("User registration details:%+v", newUser)
	if err := util.ValidatePassword(newUser.Password); err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}
	var existingUser database.User
	existingUser.Username = newUser.Username
	existingUser.Email = newUser.Email
	existingUser.PasswordHash, _ = util.HashPassword(newUser.Password)
	existingUser.CreatedAt = time.Now()
	existingUser.UpdatedAt = time.Now()

	if err := h.Repo.Create(&existingUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	util.SendData(w, existingUser, http.StatusCreated)
}
