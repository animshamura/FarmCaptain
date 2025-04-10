package controllers

import (
	"farmcaptain/models"
	"farmcaptain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

// AddCrop handles adding a new crop
func AddCrop(c *gin.Context) {
	var crop models.Crop
	if err := c.ShouldBindJSON(&crop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddCrop(crop); err != nil {
		log.Println("Error adding crop:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add crop"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Crop added successfully"})
}

// GetCrop handles fetching crop details
func GetCrop(c *gin.Context) {
	id := c.Param("id")
	crop, err := services.GetCropByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Crop not found"})
		return
	}
	c.JSON(http.StatusOK, crop)
}

// GetAIAdvice handles fetching AI-based crop guidance
func GetAIAdvice(c *gin.Context) {
	cropID := c.Param("cropId")
	advice, err := services.GetAIAdvice(cropID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get AI advice"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"advice": advice})
}
