package handlers

import "github.com/gin-gonic/gin"


func ShowLoginHandler(c *gin.Context) {
	templatePath := "login.html"
	c.HTML(200, templatePath, gin.H{
		"title": "Auth",
	})
}

func ShowRegisterHandler(c *gin.Context) {
	templatePath := "register.html"
	c.HTML(200, templatePath, gin.H{
		"title": "Auth",
	})
}

func RegisterHandler(c *gin.Context) {
	// Handle the registration form submission
}

func LoginHandler(c *gin.Context) {
	// Handle the login form submission
}

func UpdatePasswordHandler(c *gin.Context) {
	// Handle the password update form submission
}

func DeleteAccountHandler(c *gin.Context) {
	// Handle the account deletion form submission
}