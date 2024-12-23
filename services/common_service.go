package services

import "fmt"

func CreateDocumentByType(documentType string, data interface{}) (string, error) {
	fmt.Printf("Creating document of type: %s\n", documentType)

	switch documentType {
	case "document":
		fmt.Println("Processing document...")
	case "sales":
		fmt.Println("Processing sales document...")
	default:
		return "", fmt.Errorf("unsupported document type: %s", documentType)
	}

	return "new_document_id", nil
}
