package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/models"
	"github.com/nineard99/restaurant-Golang/services"
	"github.com/nineard99/restaurant-Golang/types"
)

func CreateRestaurantController(c *gin.Context) {

	var input types.RestaurantInput

	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userClaims, ok := user.(models.User)
	if !ok || strings.TrimSpace(userClaims.ID) == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
		return
	}

	input.OwnerID = userClaims.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant, err := services.CreateRestaurant(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, restaurant)
}

// GET /restaurants/:id
func GetRestaurantByIDController(c *gin.Context) {
	id := c.Param("restaurantId")

	restaurant, err := services.GetRestaurantByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

func GetAllRestaurantByUserIDController(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userClaims, ok := user.(models.User)
	if !ok || strings.TrimSpace(userClaims.ID) == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
		return
	}

	restaurants, err := services.GetAllRestaurantsByUserID(userClaims.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

func DeleteRestaurantController(c *gin.Context) {
	id := c.Param("restaurantId")

	err := services.DeleteRestaurantByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restaurant deleted successfully"})
}

func EditRestaurantNameController(c *gin.Context) {
	id := c.Param("restaurantId")

	var input struct {
		NewName string `json:"newName"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || strings.TrimSpace(input.NewName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "newName is required"})
		return
	}

	restaurant, err := services.EditRestaurantName(id, input.NewName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}
