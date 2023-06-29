package model

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}

type Product struct {
	ID          int      `json:"id"`
	SupplierID  string   `json:"supplier_id"`
	Name        string   `json:"name"`
	Image       []byte   `json:"image"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Price       float32  `json:"price"`
}

type Supplier struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
	TimeOpening string `json:"time_opening"`
	TimeClosing string `json:"time_closing"`
}
