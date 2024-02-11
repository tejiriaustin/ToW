package controllers

import (
	"context"
	"errors"
	"strings"

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
	if tokenString == "" {
		return nil, errors.New("token not set")
	}

	tokens := strings.Split(tokenString, " ")

	claims := &services.Claims{
		Content: &models.AccountInfo{},
	}

	_, err := jwt.ParseWithClaims(tokens[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return claims.Content.(*models.AccountInfo), nil
}

func IsAdmin(ctx *gin.Context, jwtSecret []byte) error {
	info, err := GetAccountInfo(ctx, jwtSecret)
	if err != nil {
		return err
	}

	if info.Kind != string(models.AdminAccount) {
		return errors.New("permission denied")
	}
	return nil
}

func GetAuthHeader(c *gin.Context) (string, error) {
	return c.GetHeader("Authorization"), nil
}
