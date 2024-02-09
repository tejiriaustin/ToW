package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/repository"
	"github.com/tejiriaustin/ToW/response"
	"github.com/tejiriaustin/ToW/services"
)

func BindRoutes(
	ctx context.Context,
	routerEngine *gin.Engine,
	sc *services.Container,
	rc *repository.Container,
	conf *env.Config,
) {

	controllers := New(ctx)

	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		response.FormatResponse(c, http.StatusOK, "OK", nil)
	})

	customer := r.Group("/customer")
	{
		customer.POST("", controllers.AccountController.CreateCustomerAccount(sc.AccountsService, sc.TokenProvider, rc.AccountsRepo))
		customer.POST("/follow/:accountId", controllers.AccountController.FollowAccount(sc.AccountsService, rc.AccountsRepo, conf))
		customer.POST("/subscribe", controllers.AccountController.Subscribe(sc.AccountsService, rc.AccountsRepo))
		customer.POST("/invest", controllers.AccountController.Invest(sc.AccountsService, rc.AccountsRepo))
		customer.POST("/buy-share", controllers.AccountController.BuyShare(sc.AccountsService, rc.AccountsRepo))
		customer.POST("/trade-wally", controllers.AccountController.TradeWally(sc.AccountsService, rc.AccountsRepo, conf))

		wallys := customer.Group("/wallys")
		wallys.POST("/transfer", nil)
		wallys.POST("/buy", nil)
	}

	admin := r.Group("/admin")
	{
		admin.POST("", controllers.AccountController.CreateAdminAccount(sc.AccountsService, sc.TokenProvider, rc.AccountsRepo))
		admin.PUT("/freeze/:accountId", controllers.AccountController.FreezeAccount(sc.AccountsService, rc.AccountsRepo))
		admin.POST("/issue-data-income", controllers.AdminController.IssueDataIncome(sc.AdminService, rc.AccountsRepo, rc.IncomeRepo, conf))
		admin.PUT("/set-minimum-follow-spend", controllers.AdminController.SetMinimumFollowSpend(sc.AdminService, conf))
	}

	//TODO: TRADE_SHARES
	// 1. buy shares
	// 2. sell shares
	// 3.

	// TODO: TRADE_CURRENCY
	// 1. buy wallys - 2% transaction fee on transfers > 100 w

	//  TODO: ADMIN FUNCTIONS
	// TODO: subscribe $10/month to access investor functions
	// TODO: function for disbursing weekly data income - manually triggered
	// 1. factor in followers
	// 2. factors in share holders - 1 Share receives 1% of total follow income
	// 3.

	// TODO: freeze accounts
	// TODO: Auto-freeze free users after X # day limit for the free version (ie 10 weeks)
	//
}
