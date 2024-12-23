package main

import (
	// "Trade_project_GO/config"
	"Trade_project_GO/routes"
	"fmt"

	// "gorm.io/gorm"
	// "fmt"
	// "log"
	// "Trade_project_GO/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// config.InitDB()
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	routes.SetupRoutes(router)
	fmt.Println("Starting server on :8080...")
	router.Run(":8080")
}
