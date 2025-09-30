package repo

import (
	"database/sql"

	"github.com/Fahedul-Islam/e-commerce/domain"
	"github.com/Fahedul-Islam/e-commerce/product"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{DB: db}
}

func (r *productRepo) Create(product *domain.Product) error {
	query := `INSERT INTO products (name, price, image_url, is_available, stock_quantity) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.DB.QueryRow(query, product.Name, product.Price, product.ImageUrl, product.IsAvailable, product.StockQuantity).Scan(&product.ID)
}

func (r *productRepo) GetAll() ([]*domain.Product, error) {
	rows, err := r.DB.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.IsAvailable, &p.StockQuantity); err != nil {
			return nil, err
		}
		products = append(products, &p) 
	}
	return products, nil
}


func (r *productRepo) GetByID(id int) (*domain.Product, error) {
	var p domain.Product
	query := `SELECT * FROM products WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.IsAvailable, &p.StockQuantity)
	return &p, err
}

func (r *productRepo) Delete(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *productRepo) Update(pId int, product *domain.Product) error {
	query := `UPDATE products SET name = $1, price = $2, image_url = $3, is_available = $4, stock_quantity = $5 WHERE id = $6`
	_, err := r.DB.Exec(query, product.Name, product.Price, product.ImageUrl, product.IsAvailable, product.StockQuantity, pId)
	return err
}
