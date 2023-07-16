package response

type SupplierResponse struct {
	ID          int                        `json:"id"`
	Categories  []SupplierCategoryResponse `json:"categories"`
	Name        string                     `json:"name"`
	ImageURL    string                     `json:"image"`
	Description string                     `json:"description"`
	TimeOpening string                     `json:"time_opening"`
	TimeClosing string                     `json:"time_closing"`
}

type SupplierCategoryResponse struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
