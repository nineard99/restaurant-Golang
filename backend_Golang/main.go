package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/models"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDB()

	// สร้างตารางอัตโนมัติ
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	r := gin.Default()

	// Controller + Route รวมกัน

	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		if err := config.DB.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูลได้"})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := config.DB.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "สร้างผู้ใช้ไม่สำเร็จ"})
			return
		}
		c.JSON(http.StatusCreated, newUser)
	})

	// เริ่ม server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
