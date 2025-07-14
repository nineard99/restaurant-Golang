package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/models"
)

func AuthorizeRestaurantRole(allowedRoles ...models.RestaurantRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAny, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}

		user := userAny.(models.User)
		restaurantID := c.Param("restaurantId")
		if restaurantID == "" {
			restaurantID = c.PostForm("restaurantId") // รองรับกรณี post
		}

		if restaurantID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing restaurantId"})
			return
		}

		var link models.RestaurantUser
		err := config.DB.Where("user_id = ? AND restaurant_id = ?", user.ID, restaurantID).First(&link).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Restaurant link not found"})
			return
		}

		for _, role := range allowedRoles {
			if link.Role == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient restaurant role"})
	}
}
