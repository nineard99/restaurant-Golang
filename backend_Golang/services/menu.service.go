package services

import (
	"errors"
	"strings"

	"github.com/nineard99/restaurant-Golang/types"
	"github.com/nineard99/restaurant-Golang/utils"

	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/models"
	"gorm.io/gorm"
)

// CreateMenuItem creates a new menu item for a restaurant
func CreateMenuItem(input types.MenuInput) (*models.MenuItem, error) {
	if strings.TrimSpace(input.Name) == "" {
		return nil, errors.New("menu name is required")
	}
	if strings.TrimSpace(input.RestaurantID) == "" {
		return nil, errors.New("restaurant ID is required")
	}

	// Check if restaurant exists
	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, "id = ?", input.RestaurantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restaurant not found")
		}
		return nil, err
	}

	// Check for duplicate menu item in same restaurant
	var existing models.MenuItem
	if err := config.DB.Where("name = ? AND restaurant_id = ?", input.Name, input.RestaurantID).First(&existing).Error; err == nil {
		return nil, errors.New("a menu item with this name already exists in this restaurant")
	}

	// Create new menu item
	menu := models.MenuItem{
		ID:           utils.GenerateUUID(),
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		RestaurantID: input.RestaurantID,
		Image:        input.Image,
	}

	if err := config.DB.Create(&menu).Error; err != nil {
		return nil, errors.New("failed to create menu item")
	}

	return &menu, nil
}

// GetAllMenuItems returns all menu items of a restaurant
func GetAllMenuItems(restaurantID string) ([]models.MenuItem, error) {
	if strings.TrimSpace(restaurantID) == "" {
		return nil, errors.New("restaurant ID is required")
	}

	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, "id = ?", restaurantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restaurant not found")
		}
		return nil, err
	}

	var menus []models.MenuItem
	err := config.DB.Where("restaurant_id = ?", restaurantID).Find(&menus).Error
	if err != nil {
		return nil, errors.New("failed to retrieve menu items")
	}

	return menus, nil
}

// GetMenuItemByID returns menu item by ID
func GetMenuItemByID(menuID string) (*models.MenuItem, error) {
	if strings.TrimSpace(menuID) == "" {
		return nil, errors.New("menu ID is required")
	}

	var menu models.MenuItem
	if err := config.DB.First(&menu, "id = ?", menuID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("menu item not found")
		}
		return nil, err
	}

	return &menu, nil
}

// DeleteMenuItem deletes a menu item
func DeleteMenuItem(menuID string) error {
	if strings.TrimSpace(menuID) == "" {
		return errors.New("menu ID is required")
	}

	var menu models.MenuItem
	if err := config.DB.First(&menu, "id = ?", menuID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("menu item not found")
		}
		return err
	}

	// TODO: Optional - delete image from cloudinary if needed

	if err := config.DB.Delete(&menu).Error; err != nil {
		return errors.New("failed to delete menu item")
	}

	return nil
}
