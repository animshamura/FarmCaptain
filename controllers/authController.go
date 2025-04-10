package controllers

import (
	"farmcaptain/models"
	"farmcaptain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login handles the farmer login and JWT token generation
func Login(c *gin.Context) {
	var farmer models.Farmer
	if err := c.ShouldBindJSON(&farmer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify farmer credentials (here we assume valid credentials)
	// In a real-world application, you would check the password from the database
	if farmer.Email == "farmer@example.com" && farmer.Password == "password" {
		token, err := services.GenerateJWT(farmer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
