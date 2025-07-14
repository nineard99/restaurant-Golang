package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/services"
	"github.com/nineard99/restaurant-Golang/types"
)

// POST /tables
func CreateSeatTableController(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	var input types.SeatTableInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.RestaurantID = restaurantId

	table, err := services.CreateSeatTable(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, table)
}

// GET
func GetAllSeatTablesController(c *gin.Context) {
	restaurantID := c.Param("restaurantId")
	tables, err := services.GetAllSeatTables(restaurantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tables)
}

// PATCH /tables/:tableId/status
func UpdateSeatTableStatusController(c *gin.Context) {
	tableID := c.Param("tableId")
	restaurantID := c.Param("restaurantId")

	var input types.SeatTableUpdateInput
	input.RestaurantID = restaurantID
	input.TableID = tableID

	if err := c.ShouldBindJSON(&input); err != nil || input.RestaurantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	table, err := services.UpdateSeatTableStatus(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, table)
}
