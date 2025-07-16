package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/controllers"
)

func OrderRoutes(r *gin.Engine) {
	order := r.Group("/:restaurantId/order")
	order.POST("/:sessionId",controllers.CreateOrderController)
	order.PATCH("/:orderId/status", controllers.UpdateOrderStatusController)
	order.GET("/:sessionId", controllers.GetAllOrderBySessionController)
	order.GET("/", controllers.GetAllOrderController)
}
