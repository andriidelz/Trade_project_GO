package repositories

import (
	"Trade_project_GO/config"
	"Trade_project_GO/models"
	"log"
)

func UpdateInventory(inventory models.Inventory) error {
	query := `UPDATE inventory SET quantity = $1 WHERE product_id = $2`
	_, err := config.DB.Exec(query, inventory.Quantity, inventory.ProductID)
	if err != nil {
		log.Printf("Error updating inventory: %v", err)
		return err
	}
	return nil
}
