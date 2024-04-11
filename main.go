package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/yemiwebby/golang-gin-hackathon-starter/routes"
)


func main() {
	router := gin.Default()

	// Initialize session middleware
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Initialize routes
	routes.InitializeRoutes(router)

	router.Run(":8080")
}