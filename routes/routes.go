package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yemiwebby/golang-gin-hackathon-starter/handlers"
)

func InitializeRoutes(router *gin.Engine) {

	// Serve static files
	router.Static("/public", "./public")
	// Load HTML temmplates
	router.LoadHTMLGlob("templates/**/*.html")

	router.GET("/", handlers.HomepageHandler)

	// Group User Routes
	// TODO: Add Middleware to this route
	userRoutes := router.Group("/user") 
	{
		// Auth routes
		userRoutes.GET("/login", handlers.ShowLoginHandler)
		userRoutes.GET("/register", handlers.ShowRegisterHandler)

		userRoutes.POST("/login", handlers.LoginHandler)
		userRoutes.POST("/register", handlers.RegisterHandler)
		userRoutes.POST("/upatePassword", handlers.UpdatePasswordHandler)
		userRoutes.POST("/deleteAccount", handlers.DeleteAccountHandler)

		// Profile route
		userRoutes.GET("/profile", handlers.ProfileHandler)
		userRoutes.POST("/profile", handlers.UpdateProfileHandler)
	}

	
}