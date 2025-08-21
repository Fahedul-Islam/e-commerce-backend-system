package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/util"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	Repo *database.AuthHandler
}

func NewUserHandler(repo *database.AuthHandler) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateTable() error {
	return h.Repo.InitTable()
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}
	if err := (newUser.Validate()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := util.ValidatePassword(newUser.Password); err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}
	var existingUser database.User
	existingUser.PasswordHash, _ = util.HashPassword(newUser.Password)
	existingUser.CreatedAt = time.Now()
	existingUser.UpdatedAt = time.Now()
	if err := h.Repo.Create(&existingUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	util.SendData(w, existingUser, http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		http.Error(w, "No users available", http.StatusNotFound)
		return
	}
	util.SendData(w, users, http.StatusOK)
}

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
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     now.Add(h.Repo.TokenExpiry).Unix(),
		"iat":     now.Unix(),
		"email":   user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.Repo.JwtSecret)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]string{"token": tokenString}, http.StatusOK)
}
