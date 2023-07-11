package response

import "time"

type SupplierResponse struct {
	ID          int                        `json:"id"`
	Categories  []SupplierCategoryResponse `json:"categories"`
	Name        string                     `json:"name"`
	Image       []byte                     `json:"image"`
	Description string                     `json:"description"`
	TimeOpening string                     `json:"time_opening"`
	TimeClosing string                     `json:"time_closing"`
}

type SupplierCategoryResponse struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type ProductResponse struct {
	ID           int      `json:"id"`
	SupplierID   int      `json:"supplier_id"`
	SupplierName string   `json:"supplier_name"`
	CategoryID   int      `json:"category_id"`
	CategoryName string   `json:"category_name"`
	Name         string   `json:"name"`
	Image        []byte   `json:"image"`
	Description  string   `json:"description"`
	Ingredients  []string `json:"ingredients"`
	Price        float32  `json:"price"`
}

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
	ProductID       int    `json:"product_id"`
	ProductName     string `json:"product_name"`
	ProductQuantity int    `json:"product_quantity"`
}
