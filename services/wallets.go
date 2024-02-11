package services

import (
	"context"
	"errors"

	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/repository"
)

type WalletService struct{}

type (
	GetWalletInput struct {
		WalletId  string
		AccountId string
	}
)

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (w WalletService) GetWallet(
	ctx context.Context,
	input GetWalletInput,
	walletRepo repository.WalletRepoInterface[*models.Wallet],
) (*models.Wallet, error) {
	if input.WalletId == "" && input.AccountId == "" {
		return nil, errors.New("an identifier is required")
	}

	filter := repository.NewQueryFilter()

	if input.WalletId != "" {
		filter.AddFilter("_id", input.WalletId)
	}
	if input.AccountId != "" {
		filter.AddFilter("account_id", input.AccountId)
	}

	return walletRepo.FindOne(ctx, filter, nil)
}

var _ WalletServiceInterface = &WalletService{}
