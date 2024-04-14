package handlers

import "github.com/gin-gonic/gin"


func ProfileHandler(c *gin.Context) {
	templatePath := "profile.html"
	c.HTML(200, templatePath, gin.H{
		"title": "Profile",
	})
}

func UpdateProfileHandler(c *gin.Context) {
	// Handle the profile form submission
}