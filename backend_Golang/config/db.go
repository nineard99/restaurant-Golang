package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nineard99/restaurant-Golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// โหลดค่า .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// อ่านค่าตัวแปรจาก env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// retry loop รอ mysql พร้อม
	var db *gorm.DB
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, errPing := db.DB()
			if errPing == nil && sqlDB.Ping() == nil {
				log.Println("✅ Connected to database")
				break
			}
		}
		log.Println("⏳ Waiting for database to be ready... retrying in 3 seconds")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect database after retries: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.RestaurantUser{},
		&models.Restaurant{},
		&models.MenuItem{},
		&models.Order{},
		&models.OrderItem{},
		models.SeatTable{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	DB = db
}
