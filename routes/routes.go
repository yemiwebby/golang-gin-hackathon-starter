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
}