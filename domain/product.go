package domain

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	ImageUrl      string  `json:"image_url"`
	IsAvailable   bool    `json:"is_available"`
	StockQuantity int     `json:"stock_quantity"`
}