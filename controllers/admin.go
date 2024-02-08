package controllers

import (
	"context"
	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/repository"
	"github.com/tejiriaustin/ToW/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/response"
)

type AdminController struct {
	conf env.Config
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (c *AdminController) IssueDataIncome(
	service services.AdminServiceInterface,
	accountsRepo repository.AccountsRepoInterface[models.Account],
	incomeRepo repository.IncomeRepoInterface[models.Income],
	config *env.Config,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		go func() {
			err := service.IssueDataIncome(context.Background(), accountsRepo, incomeRepo, config)
			if err != nil {
				response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
				return
			}
		}()
		response.FormatResponse(ctx, http.StatusOK, "successful", nil)
	}
}

func (c *AdminController) SetMinimumFollowSpend(
	adminService services.AdminServiceInterface,
	conf *env.Config,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := struct {
			MinimumFollowSpend int64 `json:"minimum_follow_spend"`
		}{}

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		_, err = GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		err = adminService.SetFollowSpend(ctx, services.SetFollowSpendInput{NewMinimumLimit: req.MinimumFollowSpend}, conf)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}
		response.FormatResponse(ctx, http.StatusOK, "successful", nil)
	}
}
