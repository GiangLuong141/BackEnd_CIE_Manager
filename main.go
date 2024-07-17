package main

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create the LoginController
	loginController := controllers.NewLoginController()

	// Create the router
	r := gin.Default()

	// Define the routes
	r.POST("/login", loginController.LoginAdmin)
	r.Run("localhost:3000")
}
