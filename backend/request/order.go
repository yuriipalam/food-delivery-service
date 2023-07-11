package request

type ProductsRequest struct {
	ProductID       int `json:"product_id"`
	ProductQuantity int `json:"product_quantity"`
}

type CreateOrderRequest struct {
	CustomerID        int               `json:"customer_id"`
	RecipientFullName string            `json:"recipient_full_name"`
	Address           string            `json:"address"`
	Price             float32           `json:"price"`
	SupplierIDs       []int             `json:"supplier_ids"`
	Products          []ProductsRequest `json:"products"`
}
