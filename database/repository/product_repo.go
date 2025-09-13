package repository

import (
	"database/sql"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}


func (r *ProductRepository) Create(product *Product) error {
	query := `INSERT INTO products (name, price, image_url, is_available, stock_quantity) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.DB.QueryRow(query, product.Name, product.Price, product.ImageUrl, product.IsAvailable, product.StockQuantity).Scan(&product.ID)
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
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.IsAvailable, &p.StockQuantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (Product, error) {
	var p Product
	query := `SELECT * FROM products WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.IsAvailable, &p.StockQuantity)
	return p, err
}

func (r *ProductRepository) Delete(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *ProductRepository) Update(pId int, product *Product) error {
	query := `UPDATE products SET name = $1, price = $2, image_url = $3, is_available = $4, stock_quantity = $5 WHERE id = $6`
	_, err := r.DB.Exec(query, product.Name, product.Price, product.ImageUrl, product.IsAvailable, product.StockQuantity, pId)
	return err
}

