package services

import (
	"Trade_project_GO/models"
	"Trade_project_GO/repositories"
	"errors"
)

func CheckStock(doc models.Document) error {
	for _, item := range doc.Items {
		product, _ := repositories.GetProductByID(int(item.ProductID))
		if product.Quantity < item.Quantity {
			return errors.New("insufficient stock for product " + product.Name)
		}
	}
	return nil
}

func ProcessFIFO(doc models.Document) {
}
