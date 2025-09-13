package repository

import (
	"errors"
	"regexp"
	"time"
)

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	ImageUrl      string  `json:"image_url"`
	IsAvailable   bool    `json:"is_available"`
	StockQuantity int     `json:"stock_quantity"`
}

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Roles        string    `json:"roles"`
}

type UserRegistration struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
}

type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Price    float64 `json:"price"`
}

type Order struct {
	OrderId   int       `json:"order_id"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	TotalPrice float64  `json:"total_price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Items     []OrderItem `json:"items,omitempty"`
}

func (o *OrderItem) Validate() error {
	if o.ProductID == 0 {
		return errors.New("product_id is required")
	}
	if o.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return nil
}

func (u *UserRegistration) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (u *UserLogin) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
