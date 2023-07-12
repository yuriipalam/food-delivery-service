package model

type Supplier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
	TimeOpening string `json:"time_opening"`
	TimeClosing string `json:"time_closing"`
}

type SupplierCategory struct {
	SupplierID int `json:"supplier_id"`
	CategoryID int `json:"category_id"`
}
