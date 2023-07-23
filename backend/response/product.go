package response

type ProductResponse struct {
	ID           int      `json:"id"`
	URL          string   `json:"url"`
	SupplierID   int      `json:"supplier_id"`
	SupplierName string   `json:"supplier_name"`
	CategoryID   int      `json:"category_id"`
	CategoryName string   `json:"category_name"`
	Name         string   `json:"name"`
	ImageURL     string   `json:"image_url"`
	Description  string   `json:"description"`
	Price        float32  `json:"price"`
}
