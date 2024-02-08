package repository

import (
	"fmt"

	"github.com/tejiriaustin/ToW/database"
	"github.com/tejiriaustin/ToW/models"
)

type Container struct {
	AccountsRepo AccountsRepoInterface[models.Account]
	WalletsRepo  WalletRepoInterface[models.Wallet]
	FollowerRepo FollowersRepoInterface[models.Follower]
	IncomeRepo   IncomeRepoInterface[models.Income]
}

var dbNameSpace = "tree_of_wally"

func New(dbConn *database.Client) *Container {
	return &Container{
		AccountsRepo: NewRepository[models.Account](dbConn.GetCollection(fmt.Sprintf("%s.tree_of_wally.accounts", dbNameSpace))),
		WalletsRepo:  NewRepository[models.Wallet](dbConn.GetCollection(fmt.Sprintf("%s.tree_of_wally.wallets", dbNameSpace))),
		IncomeRepo:   NewRepository[models.Income](dbConn.GetCollection(fmt.Sprintf("%s.tree_of_wally.follow_incomes", dbNameSpace))),
	}
}
