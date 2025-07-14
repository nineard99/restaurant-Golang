package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/routes"

	"github.com/nineard99/restaurant-Golang/utils"
)

func main() {
	config.ConnectDB()
	utils.InitCloudinary()
	r := gin.Default()

	// เพิ่ม middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ใส่ URL frontend คุณ
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // จำเป็นถ้าใช้ cookie หรือ Authorization header
		MaxAge:           12 * time.Hour,
	}))

	// routes
	routes.SetupRoutes(r)

	// route อื่นๆ...

	r.Run(":8080")
}
