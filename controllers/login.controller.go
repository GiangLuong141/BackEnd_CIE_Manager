package controllers

import (
	"example/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginController struct {
	loginService services.LoginService
}

func NewLoginController() *LoginController {
	loginService := services.NewLoginService()
	return &LoginController{
		loginService: loginService,
	}
}

func (c *LoginController) LoginAdmin(ctx *gin.Context) {
	var loginRequest LoginRequest

	token, err := c.loginService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "Logged in successfully",
	})

}
