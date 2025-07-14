package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nineard99/restaurant-Golang/controllers"
	"github.com/nineard99/restaurant-Golang/middlewares"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterController)
		auth.POST("/login", controllers.LoginController)
		auth.GET("/me", middlewares.Authenticate(), controllers.MeController)
		auth.POST("/logout", middlewares.Authenticate(), controllers.LogoutController)
	}
}
