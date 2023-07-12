package model

import "time"

type Order struct {
	ID                int       `json:"id"`
	CustomerID        int       `json:"customer_id"`
	RecipientFullName string    `json:"recipient_full_name"`
	Address           string    `json:"address"`
	Price             float32   `json:"price"`
	CreatedAt         time.Time `json:"created_at"`
}

type OrderProduct struct {
	OrderID         int `json:"order_id"`
	ProductID       int `json:"product_id"`
	ProductQuantity int `json:"product_quantity"`
}

type OrderSupplier struct {
	OrderID    int `json:"order_id"`
	SupplierID int `json:"supplier_id"`
}
