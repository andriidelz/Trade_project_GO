package routes

import (
	"Trade_project_GO/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Trade Project API!"})
	})
	router.POST("/document", controllers.CreateDocument)
	router.GET("/document", controllers.GetReport)
	router.GET("/report", controllers.GetReport)
	router.GET("/report/page", controllers.RenderReportPage)
}
