package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/services"
	"github.com/nineard99/restaurant-Golang/types"
)

func RegisterController(c *gin.Context) {

	var input types.RegisterInput

	//check input that client send Correct??
	//ShouldBindJSON return Error not True/false
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := services.RegisterUser(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เซ็ต token ลง httpOnly cookie (optional)
	c.SetCookie("jwt", token, 3600*24*7, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
		"user":  user,
	})
}

// POST /auth/login
func LoginController(c *gin.Context) {
	var input types.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := services.LoginUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// เซ็ต token ลง httpOnly cookie (optional)
	c.SetCookie("jwt", token, 3600*24*7, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func MeController(c *gin.Context) {
	user, exists := c.Get("user") // ดึงจาก context ที่ middleware ใส่ไว้

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func LogoutController(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
