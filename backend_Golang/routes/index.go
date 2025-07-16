package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	AuthRoutes(r)
	RestaurantRoutes(r)
	SeatTableRoutes(r)
	MenuRoutes(r)
	OrderRoutes(r)
	// menuRoutes(r)
	// เพิ่ม route อื่น ๆ ตามต้องการ
}
