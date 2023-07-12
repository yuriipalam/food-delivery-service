package model

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
