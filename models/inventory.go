package models

type Inventory struct {
	ProductID int     `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Cost      float64 `json:"cost" db:"cost"`
}
