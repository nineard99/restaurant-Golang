package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/controllers"
)

func MenuRoutes(r *gin.Engine) {
	menu := r.Group("/:restaurantId/menu")

	menu.POST("/", controllers.CreateMenuController)
	menu.GET("/", controllers.GetAllMenusController)

	menu.GET("/:menuId", controllers.GetMenuByIDController)
	menu.DELETE("/:menuId", controllers.DeleteMenuItemController)
}
