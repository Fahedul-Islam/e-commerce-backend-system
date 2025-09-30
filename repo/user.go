package repo

import (
	"database/sql"
	"errors"

	"github.com/Fahedul-Islam/e-commerce/domain"
	"github.com/Fahedul-Islam/e-commerce/user"
	"github.com/Fahedul-Islam/e-commerce/util"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{DB: db}
}

func (r *userRepo) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (username, email, password_hash, roles) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.DB.QueryRow(query, user.Username, user.Email, user.PasswordHash, user.Roles).Scan(&user.ID)
}

func (r *userRepo) GetAllUsers() ([]domain.User, error) {
	rows, err := r.DB.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt, &u.Roles); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepo) AuthenticateUser(login *domain.UserLogin) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE email = $1`
	if err := r.DB.QueryRow(query, login.Email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.Roles); err != nil {
		return nil, err
	}

	if err := util.CheckPasswordHash(login.Password, user.PasswordHash); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UserValidate(user *domain.UserLogin) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
