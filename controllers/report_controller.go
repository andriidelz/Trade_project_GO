package controllers

import (
	"Trade_project_GO/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	date := c.DefaultQuery("date", "")
	report, err := services.GenerateReport(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}

func RenderReportPage(c *gin.Context) {
	date := c.DefaultQuery("date", "")
	report, err := services.GenerateReport(date)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "report.tmpl", gin.H{"report": report})
}
