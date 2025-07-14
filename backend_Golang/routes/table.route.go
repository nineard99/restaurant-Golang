package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/controllers"
	"github.com/nineard99/restaurant-Golang/middlewares"
)

func SeatTableRoutes(r *gin.Engine) {
	SeatTable := r.Group("/:restaurantId/table")
	{
		SeatTable.POST("/", middlewares.Authenticate(), controllers.CreateSeatTableController)

		SeatTable.GET("/", middlewares.Authenticate(), controllers.GetAllSeatTablesController)

		SeatTable.PATCH("/:tableId",
			middlewares.Authenticate(),
			controllers.UpdateSeatTableStatusController)

	}
}
