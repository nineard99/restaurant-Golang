package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/controllers"
	"github.com/nineard99/restaurant-Golang/middlewares"
)

func RestaurantRoutes(r *gin.Engine) {
	restaurant := r.Group("/restaurant")
	{
		restaurant.POST("/create", middlewares.Authenticate(), controllers.CreateRestaurantController)

		restaurant.GET("/", middlewares.Authenticate(), controllers.GetAllRestaurantByUserIDController)

		restaurant.GET("/:restaurantId",
			middlewares.Authenticate(),
			middlewares.AuthorizeRestaurantRole("OWNER"),
			controllers.GetRestaurantByIDController)

		restaurant.PATCH("/:restaurantId/editName",
			middlewares.Authenticate(),
			middlewares.AuthorizeRestaurantRole("OWNER"),
			controllers.EditRestaurantNameController)

		restaurant.DELETE("/:restaurantId",
			middlewares.Authenticate(),
			controllers.DeleteRestaurantController)
	}
}
