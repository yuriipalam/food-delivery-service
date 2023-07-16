package response

import "time"

type OrderResponse struct {
	ID                int                     `json:"id"`
	CustomerID        int                     `json:"customer_id"`
	CustomerName      string                  `json:"customer_name"`
	Suppliers         []OrderSupplierResponse `json:"suppliers"`
	Products          []OrderProductResponse  `json:"products"`
	RecipientFullName string                  `json:"recipient_full_name"`
	Address           string                  `json:"address"`
	Price             float32                 `json:"price"`
	CreatedAt         time.Time               `json:"created_at"`
}

type OrderSupplierResponse struct {
	SupplierID   int    `json:"supplier_id"`
	SupplierName string `json:"supplier_name"`
}

type OrderProductResponse struct {
	ProductID         int    `json:"product_id"`
	ProductSupplierID int    `json:"product_supplier_id"`
	ProductName       string `json:"product_name"`
	ProductQuantity   int    `json:"product_quantity"`
}
