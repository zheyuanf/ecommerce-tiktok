package model

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Picture     string  `json:"picture"`
}
