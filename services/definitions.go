package services

import (
	"context"

	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/repository"
)

type AccountServiceInterface interface {
	CreateAccount(
		ctx context.Context,
		input CreateAccountInput,
		accountsRepo repository.AccountsRepoInterface[*models.Account],
	) (*models.Account, error)

	GetAccount(ctx context.Context,
		input GetAccountInput,
		accountsRepo repository.AccountsRepoInterface[*models.Account],
	) (*models.Account, error)

	FreezeAccount(
		ctx context.Context,
		input FreezeAccountInput,
		accountsRepo repository.AccountsRepoInterface[*models.Account],
	) (*models.Account, error)

	Subscribe(
		ctx context.Context,
		input SubscribeAccountInput,
		accountsRepo repository.AccountsRepoInterface[*models.Account],
	) (*models.Account, error)

	BuyShare(
		ctx context.Context,
		input InvestAccountInput,
		accountsRepo repository.AccountsRepoInterface[*models.Account],
		walletsRepo repository.WalletRepoInterface[*models.Wallet],
		conf *env.Config,
	) error

	FollowAccount(
		ctx context.Context,
		input FollowAccountInput,
		accountsRepo repository.AccountsRepoInterface[*models.Account],
		followerRepo repository.FollowersRepoInterface[*models.Follower],
		config *env.Config,
	) error

	TradeWally(
		ctx context.Context,
		input TradeWallyInput,
		accountsRepo repository.AccountsRepoInterface[*models.Account],
		config *env.Config,
	) error
}

type AdminServiceInterface interface {
	SetFollowSpend(
		ctx context.Context,
		input SetFollowSpendInput,
		config *env.Config,
	) error

	IssueDataIncome(
		ctx context.Context,
		finder repository.FindCursor,
		incomeRepo repository.IncomeRepoInterface[*models.Income],
	) error
}

type WalletServiceInterface interface {
	GetWallet(
		ctx context.Context,
		input GetWalletInput,
		walletRepo repository.WalletRepoInterface[*models.Wallet],
	) (*models.Wallet, error)
}
