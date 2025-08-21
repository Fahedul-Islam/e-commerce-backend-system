package database

import (
	"database/sql"
	"log"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) InitTable() error {
	query := `CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price DECIMAL(10,2) NOT NULL,
		image_url VARCHAR(255) NOT NULL
	)`
	log.Printf("Executing query: %s", query)
	_, err := r.DB.Exec(query)
	return err
}

func (r *ProductRepository) Create(product *Product) error {
	query := `INSERT INTO products (name, price, image_url) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, product.Name, product.Price, product.ImageUrl).Scan(&product.ID)
}

func (r *ProductRepository) GetAll() ([]Product, error) {
	rows, err := r.DB.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (Product, error) {
	var p Product
	query := `SELECT * FROM products WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl)
	return p, err
}
