package response

type ProductResponse struct {
	ID           int      `json:"id"`
	SupplierID   int      `json:"supplier_id"`
	SupplierName string   `json:"supplier_name"`
	CategoryID   int      `json:"category_id"`
	CategoryName string   `json:"category_name"`
	Name         string   `json:"name"`
	ImageURL     string   `json:"image"`
	Description  string   `json:"description"`
	Ingredients  []string `json:"ingredients"`
	Price        float32  `json:"price"`
}
