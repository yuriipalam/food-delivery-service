package model

import "time"

type Customer struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}

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
