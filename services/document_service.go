package services

import (
	"Trade_project_GO/models"
	"Trade_project_GO/repositories"
	"Trade_project_GO/utils"
)

func CreateDocument(doc models.Document) error {
	if err := utils.ValidateDocument(doc); err != nil {
		return err
	}

	if doc.DocumentType == "outcome" {
		inventory, err := repositories.GetInventoryByProductID(int(doc.ID))
		if err != nil {
			return err
		}

		updatedInventory, err := utils.FIFOProcessStock([]models.Inventory{*inventory}, int(doc.ID), int(doc.TotalAmount))
		if err != nil {
			return err
		}

		for _, inv := range updatedInventory {
			err = repositories.UpdateInventory(inv)
			if err != nil {
				return err
			}
		}
	}
	return repositories.SaveDocument(doc)
}

// 		err = repositories.UpdateInventory(config.DB, &updatedInventory)

// 		// err = repositories.UpdateInventory(updatedInventory)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return repositories.SaveDocument(doc)
// }
