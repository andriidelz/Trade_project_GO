package repositories

import (
	"Trade_project_GO/config"
	"Trade_project_GO/models"

	// "Trade_project_GO/repositories"
	"log"

	"gorm.io/gorm"
)

func CreateDocument(db *gorm.DB, doc *models.Document) error {
	if err := db.Create(doc).Error; err != nil {
		return err
	}
	return nil
}

func SaveDocument(doc models.Document) error {
	query := `INSERT INTO documents (document_type, date, total_amount) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, doc.DocumentType, doc.Date, doc.TotalAmount).Scan(&doc.ID)
	// _, err := config.DB.Exec(query, doc.DocumentType, doc.Date, doc.TotalAmount)
	if err != nil {
		log.Printf("Error saving document: %v", err)
		return err
	}
	return nil
}

func GetDocumentsByDate(startDate, endDate string) ([]models.Document, error) {
	var documents []models.Document
	query := `SELECT * FROM documents WHERE date BETWEEN $1 AND $2`
	err := config.DB.Select(&documents, query, startDate, endDate)
	if err != nil {
		log.Printf("Error fetching documents: %v", err)
		return nil, err
	}
	return documents, nil
}

func UpdateDocument(doc models.Document) error {
	query := `UPDATE documents SET document_type = $1, date = $2, total_amount = $3 WHERE id = $4`
	_, err := config.DB.Exec(query, doc.DocumentType, doc.Date, doc.TotalAmount, doc.ID) // Fixed missing parameter
	return err
}

func GetDocumentByID(id int) (models.Document, error) {
	var doc models.Document
	err := config.DB.Get(&doc, "SELECT * FROM documents WHERE id = $1", id)
	return doc, err
}

func GetInventoryByProductID(productID int) (*models.Inventory, error) {
	var inventory models.Inventory
	// query := "SELECT * FROM inventory WHERE product_id = $1"
	err := config.DB.Get(&inventory, "SELECT * FROM inventory WHERE product_id = $1", productID)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}
