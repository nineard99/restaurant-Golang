package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/controllers"
	"github.com/nineard99/restaurant-Golang/middlewares"
)

func MenuRoutes(r *gin.Engine) {
	menu := r.Group("/:restaurantId/menu")

	menu.POST("/", middlewares.Authenticate(), controllers.CreateMenuController)
	menu.GET("/", middlewares.Authenticate(), controllers.GetAllMenusController)
	menu.GET("/:menuId", middlewares.Authenticate(), controllers.GetMenuByIDController)
	menu.DELETE("/:menuId", middlewares.Authenticate(), controllers.DeleteMenuItemController)
}
