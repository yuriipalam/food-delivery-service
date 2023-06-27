package model

type Product struct {
	ID          int      `json:"id"`
	SupplierID  string   `json:"supplier_id"`
	Name        string   `json:"name"`
	Image       []byte   `json:"image"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Price       float32  `json:"price"`
}
