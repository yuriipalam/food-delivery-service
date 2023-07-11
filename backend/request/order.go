package request

type CreateOrder struct {
	CustomerID             int         `json:"customer_id"`
	RecipientFullName      string      `json:"recipient_full_name"`
	Address                string      `json:"address"`
	Price                  float32     `json:"price"`
	SupplierIDs            []int       `json:"supplier_ids"`
	ProductIDQuantityPairs map[int]int `json:"product_id_quantity_pairs"`
}
