package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/services"
	"github.com/nineard99/restaurant-Golang/types"

	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/utils"
)

func CreateMenuController(c *gin.Context) {
	restaurantId := c.Param("restaurantId")

	var input types.MenuInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	input.RestaurantID = restaurantId

	// สร้างเมนูก่อน (ยังไม่ใส่รูป)
	menu, err := services.CreateMenuItem(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ถ้ามีไฟล์ภาพแนบมา
	fileHeader, err := c.FormFile("image")
	if err == nil && fileHeader != nil {
		file, _ := fileHeader.Open()
		defer file.Close()

		imageUrl, err := utils.UploadImageToCloudinary(file, "menus", menu.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}

		// อัปเดตเมนูให้มีรูป
		menu.Image = &imageUrl
		config.DB.Model(menu).Update("image", imageUrl)
	}

	c.JSON(http.StatusCreated, menu)
}

// GET /menus/:restaurantId
func GetAllMenusController(c *gin.Context) {
	restaurantId := c.Param("restaurantId")

	menus, err := services.GetAllMenuItems(restaurantId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func GetMenuByIDController(c *gin.Context) {
	menuId := c.Param("menuId")

	menu, err := services.GetMenuItemByID(menuId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

// DELETE /menu/:menuId
func DeleteMenuItemController(c *gin.Context) {
	menuId := c.Param("menuId")

	err := services.DeleteMenuItem(menuId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu item deleted successfully."})
}
