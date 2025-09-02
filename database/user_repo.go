package database

import (
	"database/sql"
	"time"

	"github.com/Fahedul-Islam/e-commerce/util"
)

type AuthHandler struct {
	DB          *sql.DB
	JwtSecret   []byte
	TokenExpiry time.Duration
}

func NewAuthHandler(db *sql.DB, jwtSecret []byte) *AuthHandler {
	return &AuthHandler{DB: db, JwtSecret: jwtSecret, TokenExpiry: 24 * time.Hour}
}

func (r *AuthHandler) Create(user *User) error {
	query := `INSERT INTO users (username, email, password_hash, roles) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.DB.QueryRow(query, user.Username, user.Email, user.PasswordHash, user.Roles).Scan(&user.ID)
}

func (r *AuthHandler) GetAll() ([]User, error) {
	rows, err := r.DB.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt, &u.Roles); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *AuthHandler) Authenticate(email, password string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email = $1`
	if err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.Roles); err != nil {
		return nil, err
	}

	if err := util.CheckPasswordHash(password, user.PasswordHash); err != nil {
		return nil, err
	}
	return &user, nil
}
