package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Fahedul-Islam/e-commerce/config"
	"github.com/Fahedul-Islam/e-commerce/database/connections"
	"github.com/Fahedul-Islam/e-commerce/domain"
	"github.com/Fahedul-Islam/e-commerce/util"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	cnf  *config.Config
	srvc Service
}

type contextKey string

const (
	ContextUserID contextKey = "user_id"
	ContextEmail  contextKey = "email"
	ContextRoles  contextKey = "roles"
)

func NewUserHandler(cnf *config.Config, srvc Service) *UserHandler {
	return &UserHandler{cnf: cnf,
		srvc: srvc,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.srvc.GetAllUsers()
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

func (h *UserHandler) generateToken(user *domain.User) (string, string, error) {
	now := time.Now()
	accessClaims := jwt.MapClaims{
		"user_id": strconv.Itoa(int(user.ID)),
		"exp":     now.Add(h.cnf.JWT.TokenExpiry).Unix(),
		"iat":     now.Unix(),
		"email":   user.Email,
		"roles":   user.Roles,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(h.cnf.JWT.Secret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.MapClaims{
		"user_id": strconv.Itoa(int(user.ID)),
		"exp":     now.Add(h.cnf.JWT.RefreshExpiry).Unix(),
		"iat":     now.Unix(),
		"email":   user.Email,
		"roles":   user.Roles,
	}

	refreshJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken)
	refreshTokenString, err := refreshJwt.SignedString(h.cnf.JWT.Secret)
	if err != nil {
		return "", "", err
	}

	key := fmt.Sprintf("refresh_token_%d", user.ID)
	if err := connections.SetRedisClient(key, refreshTokenString, h.cnf.JWT.RefreshExpiry); err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (h *UserHandler) GenerateRefreshToken(w http.ResponseWriter, r *http.Request) (string, string, error) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return "", "", errors.New("invalid request body")
	}

	claim := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(body.RefreshToken, claim, func(token *jwt.Token) (interface{}, error) {
		return h.cnf.JWT.Secret, nil
	})
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	userID := claim["user_id"].(float64)
	key := fmt.Sprintf("refresh_token_%d", int(userID))
	storedToken, err := connections.GetRedisClient(key)
	if err != nil || storedToken != body.RefreshToken {
		return "", "", errors.New("refresh token not recognized")
	}
	user := &domain.User{
		ID:    int(userID),
		Email: claim["email"].(string),
		Roles: claim["roles"].(string),
	}
	newAccessToken, newRefreshToken, err := h.generateToken(user)
	if err != nil {
		return "", "", errors.New("failed to generate new tokens")
	}

	return newAccessToken, newRefreshToken, nil
}
