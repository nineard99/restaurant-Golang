package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/models"
	"github.com/nineard99/restaurant-Golang/types"
)

func Authenticate() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		// 1. ดึง token จาก cookie หรือ header
		tokenString, err := getTokenFromRequest(c)
		if err != nil || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token not provided"})
			return
		}

		// 2. แปลง JWT
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "default_secret"
		}

		token, err := jwt.ParseWithClaims(tokenString, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid token"})
			return
		}

		// 3. ดึง claims และหา user
		claims, ok := token.Claims.(*types.JWTClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid claims"})
			return
		}

		var user models.User
		if err := config.DB.First(&user, "id = ?", claims.ID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: user not found"})
			return
		}

		// 4. แนบ user กับ context
		c.Set("user", user)
		c.Next()
	}
}

func getTokenFromRequest(c *gin.Context) (string, error) {

	if token, err := c.Cookie("jwt"); err == nil {
		return token, nil
	}

	// จาก header Authorization: Bearer
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), nil
	}

	return "", nil
}
