package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/services"
	"github.com/nineard99/restaurant-Golang/types"
)

func CreateOrderController(c *gin.Context) {
	var input types.CreateOrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurantId := c.Param("restaurantId")
	sessionId := c.Param("sessionId")
	input.RestaurantID = restaurantId
	input.SessionID = sessionId
	newOrder, err := services.CreateOrder(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

func GetAllOrderBySessionController(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	sessionId := c.Param("sessionId")

	orders, err := services.GetOrdersBySession(restaurantId, sessionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetAllOrderController(c *gin.Context) {
	restaurantId := c.Param("restaurantId")

	orders, err := services.GetAllOrders(restaurantId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func UpdateOrderStatusController(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	orderId := c.Param("orderId")

	var input struct {
		NewStatus string `json:"newStatus"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := services.UpdateOrderStatus(restaurantId, orderId, input.NewStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
