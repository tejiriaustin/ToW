package controllers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/services"
)

type Controller struct {
	AccountController *AccountController
	AdminController   *AdminController
}

func New(ctx context.Context) *Controller {
	return &Controller{
		AccountController: NewAccountController(),
	}
}

func GetAccountInfo(ctx *gin.Context, jwtSecret []byte) (*models.AccountInfo, error) {
	tokenString, _ := GetAuthHeader(ctx)
	if tokenString != "" {
		return nil, errors.New("token not set")
	}

	claims := &services.Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return claims.Content.(*models.AccountInfo), nil
}

func GetAuthHeader(c *gin.Context) (string, error) {
	return c.GetHeader("x-token-user"), nil
}
