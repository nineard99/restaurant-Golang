package services

import (
	"errors"
	"strings"

	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/models"
	"github.com/nineard99/restaurant-Golang/types"
	"github.com/nineard99/restaurant-Golang/utils"
	"gorm.io/gorm"
)

func CreateSeatTable(input types.SeatTableInput) (*models.SeatTable, error) {
	if strings.TrimSpace(input.Name) == "" {
		return nil, errors.New("table name is required")
	}
	if strings.TrimSpace(input.RestaurantID) == "" {
		return nil, errors.New("restaurant ID is required")
	}

	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, "id = ?", input.RestaurantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restaurant not found")
		}
		return nil, err
	}

	var existing models.SeatTable
	if err := config.DB.Where("name = ? AND restaurant_id = ?", input.Name, input.RestaurantID).First(&existing).Error; err == nil {
		return nil, errors.New("table name already exists")
	}

	table := models.SeatTable{
		ID:           utils.GenerateUUID(),
		Name:         input.Name,
		IsActive:     false,
		RestaurantID: input.RestaurantID,
	}

	if err := config.DB.Create(&table).Error; err != nil {
		return nil, errors.New("failed to create table")
	}

	return &table, nil
}
func GetAllSeatTables(restaurantID string) ([]models.SeatTable, error) {
	if strings.TrimSpace(restaurantID) == "" {
		return nil, errors.New("restaurant ID is required")
	}

	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, "id = ?", restaurantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ร้านไม่เจอ -> return array ว่างเลย (หรือจะ return error ก็ได้)
			return []models.SeatTable{}, nil
		}
		return nil, err
	}

	var tables []models.SeatTable
	if err := config.DB.Where("restaurant_id = ?", restaurantID).Find(&tables).Error; err != nil {
		// ดึงโต๊ะไม่สำเร็จ ให้ return array ว่างแทน error
		return []models.SeatTable{}, nil
	}

	// ถึงจะไม่มีโต๊ะเลย tables ก็เป็น slice ว่างอยู่แล้ว
	return tables, nil
}

func UpdateSeatTableStatus(input types.SeatTableUpdateInput) (*models.SeatTable, error) {
	var table models.SeatTable
	if err := config.DB.First(&table, "id = ? AND restaurant_id = ?", input.TableID, input.RestaurantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("table not found in this restaurant")
		}
		return nil, err
	}

	var sessionID *string
	if input.IsActive {
		id := utils.GenerateUUID()
		sessionID = &id
	}

	table.IsActive = input.IsActive
	table.SessionID = sessionID

	if input.CurrentOccupancy != nil {
		table.CurrentOccupancy = input.CurrentOccupancy
	}

	if err := config.DB.Save(&table).Error; err != nil {
		return nil, errors.New("failed to update table status")
	}

	return &table, nil
}
