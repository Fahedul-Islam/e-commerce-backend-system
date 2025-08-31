package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/util"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	Repo *database.AuthHandler
}

type contextKey string

const (
	ContextUserID contextKey = "user_id"
	ContextEmail  contextKey = "email"
)

func NewUserHandler(repo *database.AuthHandler) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateTable() error {
	return h.Repo.InitTable()
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

func (h *UserHandler) generateToken(user *database.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     now.Add(h.Repo.TokenExpiry).Unix(),
		"iat":     now.Unix(),
		"email":   user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.Repo.JwtSecret)
}

func (h *UserHandler) GenerateRefreshToken(w http.ResponseWriter, r *http.Request) (string, error) {
	ctx := r.Context()
	userId := ctx.Value(ContextUserID)
	email := ctx.Value(ContextEmail)

	if userId == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return "", errors.New("unauthorized")
	}
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     now.Add(h.Repo.TokenExpiry).Unix(),
		"iat":     now.Unix(),
		"email":   email.(string),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.Repo.JwtSecret)
}
