package repository

import (
	"github.com/tejiriaustin/ToW/database"
	"github.com/tejiriaustin/ToW/models"
)

type Container struct {
	AccountsRepo AccountsRepoInterface[*models.Account]
	WalletsRepo  WalletRepoInterface[*models.Wallet]
	FollowerRepo FollowersRepoInterface[*models.Follower]
	IncomeRepo   IncomeRepoInterface[*models.Income]
}

func New(dbConn *database.Client) *Container {
	return &Container{
		AccountsRepo: NewRepository[*models.Account](dbConn.GetCollection("accounts")),
		WalletsRepo:  NewRepository[*models.Wallet](dbConn.GetCollection("wallets")),
		IncomeRepo:   NewRepository[*models.Income](dbConn.GetCollection("incomes")),
		FollowerRepo: NewRepository[*models.Follower](dbConn.GetCollection("followers")),
	}
}
