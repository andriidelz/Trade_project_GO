package utils

import (
	"Trade_project_GO/models"
	"errors"
)

func FIFOProcessStock(inventory []models.Inventory, productID int, quantity int) ([]models.Inventory, error) {
	var remainingQuantity = quantity
	for i, inv := range inventory {
		if inv.ProductID == productID && inv.Quantity > 0 {
			if inv.Quantity >= remainingQuantity {
				inventory[i].Quantity -= remainingQuantity
				return inventory, nil
			} else {
				remainingQuantity -= inv.Quantity
				inventory[i].Quantity = 0
			}
		}
	}
	return inventory, errors.New("insufficient stock to fulfill the request")
}
