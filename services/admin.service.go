package services

import (
	"errors"
	"example/responses"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginService interface {
	Login(email, password string) (string, error)
}

type loginService struct {
	adminRepository responses.AdminRepository
	jwtSecret       string
}

func NewLoginService() LoginService {
	return &loginService{}
}

func (s *loginService) Login(email, password string) (string, error) {
	user, err := s.adminRepository.GetUserByEmail(email)
	if err != nil {
		return "invalid email or password ", err
	}

	fmt.Println(user)

	if user.Password != password {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
