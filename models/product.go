package models

type Product struct {
    ID       int     `json:"id" db:"id"`
    Name     string  `json:"name" db:"name"`
    Quantity int     `json:"quantity" db:"quantity"`
    Price    float64 `json:"price" db:"price"`
}
