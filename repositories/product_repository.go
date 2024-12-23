package repositories

import (
	"Trade_project_GO/config"
	"Trade_project_GO/models"
)

func GetProductByID(id int) (models.Product, error) {
	var product models.Product
	err := config.DB.Get(&product, "SELECT * FROM products WHERE id = $1", id)
	return product, err
}

func UpdateProductQuantity(id, quantity int) error {
	_, err := config.DB.Exec("UPDATE products SET quantity = quantity - $1 WHERE id = $2", quantity, id)
	return err
}
