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

func CreateRestaurant(input types.RestaurantInput) (*models.Restaurant, error) {
	if strings.TrimSpace(input.Name) == "" {
		return nil, errors.New("restaurant name is required")
	}
	if strings.TrimSpace(input.OwnerID) == "" {
		return nil, errors.New("owner ID is required")
	}

	var existing models.Restaurant

	if err := config.DB.Where("name = ?", input.Name).First(&existing).Error; err == nil {
		return nil, errors.New("a restaurant with this name already exists")
	}

	restaurantID := utils.GenerateUUID()

	restaurant := models.Restaurant{
		ID:      restaurantID,
		Name:    input.Name,
		Address: input.Address,
		Image:   input.Image,
		Users: []models.RestaurantUser{
			{
				ID:     utils.GenerateUUID(),
				UserID: input.OwnerID,
				Role:   models.RestaurantRoleOwner,
			},
		},
	}

	if err := config.DB.Create(&restaurant).Error; err != nil {
		return nil, errors.New("failed to create restaurant")
	}

	return &restaurant, nil
}

func GetRestaurantByID(restaurantID string) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, "id = ?", restaurantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restaurant not found")
		}
		return nil, errors.New("error retrieving restaurant")
	}
	return &restaurant, nil
}

func GetAllRestaurantsByUserID(userID string) ([]map[string]interface{}, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, errors.New("user ID is required")
	}

	var userRestaurants []models.RestaurantUser
	if err := config.DB.Preload("Restaurant").Where("user_id = ?", userID).Find(&userRestaurants).Error; err != nil {
		return nil, errors.New("failed to retrieve user's restaurants")
	}

	var result []map[string]interface{}
	for _, ur := range userRestaurants {
		result = append(result, map[string]interface{}{
			"id":        ur.Restaurant.ID,
			"name":      ur.Restaurant.Name,
			"createdAt": ur.Restaurant.CreatedAt,
			"role":      ur.Role,
		})
	}

	return result, nil
}
func DeleteRestaurantByID(restaurantID string) error {
	if strings.TrimSpace(restaurantID) == "" {
		return errors.New("restaurant ID is required")
	}

	return config.DB.Transaction(func(tx *gorm.DB) error {
		var restaurant models.Restaurant
		if err := tx.First(&restaurant, "id = ?", restaurantID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("restaurant not found")
			}
			return errors.New("failed to find restaurant")
		}

		// ลบ order items ผ่าน join orders ก่อน
		var orders []models.Order
		if err := tx.Where("restaurant_id = ?", restaurantID).Find(&orders).Error; err != nil {
			return err
		}

		orderIDs := make([]string, len(orders))
		for i, order := range orders {
			orderIDs[i] = order.ID
		}

		if len(orderIDs) > 0 {
			if err := tx.Where("order_id IN ?", orderIDs).Delete(&models.OrderItem{}).Error; err != nil {
				return err
			}
		}

		// ลบ orders ที่เป็นของร้านนี้
		if err := tx.Where("restaurant_id = ?", restaurantID).Delete(&models.Order{}).Error; err != nil {
			return err
		}

		// ลบเมนูที่เป็นของร้านนี้
		if err := tx.Where("restaurant_id = ?", restaurantID).Delete(&models.MenuItem{}).Error; err != nil {
			return err
		}

		// ลบโต๊ะที่เป็นของร้านนี้
		if err := tx.Where("restaurant_id = ?", restaurantID).Delete(&models.SeatTable{}).Error; err != nil {
			return err
		}

		// ลบความสัมพันธ์กับผู้ใช้
		if err := tx.Where("restaurant_id = ?", restaurantID).Delete(&models.RestaurantUser{}).Error; err != nil {
			return err
		}

		// ลบร้านอาหาร
		if err := tx.Delete(&restaurant).Error; err != nil {
			return err
		}

		return nil
	})
}

func EditRestaurantName(restaurantID, newName string) (*models.Restaurant, error) {
	if strings.TrimSpace(restaurantID) == "" || strings.TrimSpace(newName) == "" {
		return nil, errors.New("restaurant ID and new name are required")
	}

	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, "id = ?", restaurantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restaurant not found")
		}
		return nil, errors.New("failed to retrieve restaurant")
	}

	var nameCheck models.Restaurant
	if err := config.DB.Where("name = ? AND id != ?", newName, restaurantID).First(&nameCheck).Error; err == nil {
		return nil, errors.New("a restaurant with this name already exists")
	}

	restaurant.Name = newName
	if err := config.DB.Save(&restaurant).Error; err != nil {
		return nil, errors.New("failed to update restaurant name")
	}

	return &restaurant, nil
}
