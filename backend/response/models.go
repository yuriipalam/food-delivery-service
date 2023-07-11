package response

import "time"

type SupplierResponse struct {
	ID            int      `json:"id"`
	CategoryIDs   []int    `json:"category_ids"`
	CategoryNames []string `json:"category_names"`
	Name          string   `json:"name"`
	Image         []byte   `json:"image"`
	Description   string   `json:"description"`
	TimeOpening   string   `json:"time_opening"`
	TimeClosing   string   `json:"time_closing"`
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
	ID                int       `json:"id"`
	CustomerID        int       `json:"customer_id"`
	CustomerName      string    `json:"customer_name"`
	SupplierIDs       []int     `json:"supplier_ids"`
	SupplierNames     []string  `json:"supplier_names"`
	ProductIDs        []int     `json:"product_ids"`
	ProductName       []string  `json:"product_names"`
	RecipientFullName string    `json:"recipient_full_name"`
	Address           string    `json:"address"`
	Price             float32   `json:"price"`
	CreatedAt         time.Time `json:"created_at"`
}
