package model

type Product struct {
	ID          int      `json:"id"`
	SupplierID  int      `json:"supplier_id"`
	CategoryID  int      `json:"category_id"`
	Name        string   `json:"name"`
	Image       []byte   `json:"image"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Price       float32  `json:"price"`
}
