package models

type Document struct {
	ID           uint           `gorm:"primaryKey"`
	Date         string         `gorm:"not null"`
	TotalAmount  float64        `gorm:"not null"`
	DocumentType string         `gorm:"not null"`
	Items        []DocumentItem `gorm:"foreignKey:DocumentID"`
}

type DocumentItem struct {
	ID          uint    `gorm:"primaryKey"`
	DocumentID  uint    `gorm:"not null"`
	ProductID   uint    `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	UnitPrice   float64 `gorm:"not null"`
	TotalAmount float64 `gorm:"not null"`
}

// import "time"

// type Document struct {
//     ID          int       `json:"id" db:"id"`
//     DocumentType string  `json:"document_type" db:"document_type"`
//     Date        time.Time `json:"date" db:"date"`
//     TotalAmount float64  `json:"total_amount" db:"total_amount"`
// }
