package handlers

import "github.com/gin-gonic/gin"


func HomepageHandler(c *gin.Context) {
	templatePath := "index.html"
	c.HTML(200, templatePath, gin.H{
		"title": "Golang Gin Hackathon Starter",
	})
}