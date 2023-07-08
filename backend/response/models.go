package response

type SupplierResponse struct {
	ID           int    `json:"id"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Name         string `json:"name"`
	Image        []byte `json:"image"`
	Description  string `json:"description"`
	TimeOpening  string `json:"time_opening"`
	TimeClosing  string `json:"time_closing"`
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
