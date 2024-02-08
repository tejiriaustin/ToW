package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/repository"
	"github.com/tejiriaustin/ToW/requests"
	"github.com/tejiriaustin/ToW/response"
	"github.com/tejiriaustin/ToW/services"
)

type AccountController struct {
	conf env.Config
}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (c *AccountController) CreateCustomerAccount(
	accountService services.AccountServiceInterface,
	tokenProvider services.TokenProvider,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		input := services.CreateAccountInput{
			FirstName:     req.FirstName,
			LastName:      req.LastName,
			Phone:         req.Phone,
			DOB:           time.Time{},
			Country:       req.Country,
			ZipCode:       req.ZipCode,
			Email:         req.Email,
			Profession:    req.Profession,
			Income:        req.Income,
			Company:       req.Company,
			PersonalLinks: req.PersonalLinks,
			Kind:          models.CustomerAccount,
		}

		account, err := accountService.CreateAccount(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		account.Token, err = tokenProvider.GenerateToken(account)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(&account))
	}
}

func (c *AccountController) CreateAdminAccount(
	accountService services.AccountServiceInterface,
	tokenProvider services.TokenProvider,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		input := services.CreateAccountInput{
			FirstName:     req.FirstName,
			LastName:      req.LastName,
			Phone:         req.Phone,
			DOB:           time.Time{},
			Country:       req.Country,
			ZipCode:       req.ZipCode,
			Email:         req.Email,
			Profession:    req.Profession,
			Income:        req.Income,
			Company:       req.Company,
			PersonalLinks: req.PersonalLinks,
			Kind:          models.AdminAccount,
		}

		account, err := accountService.CreateAccount(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		account.Token, err = tokenProvider.GenerateToken(account)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(&account))
	}
}

func (c *AccountController) FreezeAccount(
	accountService services.AccountServiceInterface,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accountId := ctx.Param("accountId")

		_, err := accountService.FreezeAccount(ctx, services.FreezeAccountInput{AccountId: accountId}, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", nil)
	}
}

func (c *AccountController) FollowAccount(
	accountService services.AccountServiceInterface,
	accountsRepo repository.AccountsRepoInterface[models.Account],
	conf *env.Config,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		account, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}
		input := services.FollowAccountInput{
			AccountId:         ctx.Param("accountId"),
			FollowerAccountId: account.Id,
		}

		err = accountService.FollowAccount(ctx, input, accountsRepo, conf)
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
			return
		}
		response.FormatResponse(ctx, http.StatusOK, "successful", nil)
	}
}

func (c *AccountController) Subscribe(
	accountService services.AccountServiceInterface,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accountInfo, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		account, err := accountService.Subscribe(ctx, services.SubscribeAccountInput{AccountId: accountInfo.Id}, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(account))

	}
}

func (c *AccountController) Invest(
	accountService services.AccountServiceInterface,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", nil)

	}
}

func (c *AccountController) BuyShare(
	accountService services.AccountServiceInterface,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", nil)

	}
}
func (c *AccountController) TradeWally(
	accountService services.AccountServiceInterface,
	accountsRepo repository.AccountsRepoInterface[models.Account],
	conf *env.Config,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.TradeWallyRequest
		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "bad request", nil)
			return
		}

		_, err = GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		input := services.TradeWallyInput{
			Amount:           req.Amount,
			RecipientDetails: req.RecipientDetails,
		}
		err = accountService.TradeWallys(ctx, input, accountsRepo, conf)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", nil)
	}
}
