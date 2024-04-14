package handlers

import "github.com/gin-gonic/gin"


func ShowContactHandler(c *gin.Context) {
	templatePath := "contact.html"
	c.HTML(200, templatePath, gin.H{
		"title": "Contact",
	})
}


func ContactHandler(c *gin.Context) {
	// Handle the contact form submission
}