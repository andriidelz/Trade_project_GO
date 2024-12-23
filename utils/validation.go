package utils

import (
	"Trade_project_GO/models"
	"errors"
	"time"
)

func ValidateProductQuantity(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	return nil
}

func ValidateDocument(doc models.Document) error {
	if doc.DocumentType != "income" && doc.DocumentType != "outcome" {
		return errors.New("invalid document type")
	}
	if doc.TotalAmount <= 0 {
		return errors.New("total amount must be greater than zero")
	}
	return nil
}

func ValidateDate(dateStr string) (bool, error) {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

// func ValidateDate(date string) (string, error) {
// 	_, err := strconv.ParseTime(date, "2006-01-02")
// 	if err != nil {
// 		return "", errors.New("invalid date format, expected YYYY-MM-DD")
// 	}
// 	return date, nil
// }
