package services

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/repository"
)

type AdminService struct{}

var (
	defaultTotalDataIncome int64 = 100
)

type SetFollowSpendInput struct {
	NewMinimumLimit int64
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (w AdminService) SetFollowSpend(
	ctx context.Context,
	input SetFollowSpendInput,
	config *env.Config,
) error {
	if input.NewMinimumLimit <= 0 {
		return errors.New("minimum follow spend limit cannot be zero or negative")
	}

	config.SetEnv(env.MinimumFollowSpend, input.NewMinimumLimit)
	return nil
}

func (w AdminService) IssueDataIncome(
	ctx context.Context,
	finder repository.FindCursor,
	incomeRepo repository.IncomeRepoInterface[*models.Income],
) error {

	filters := repository.NewQueryFilter()
	accountGenerator, err := repository.NewModelGenerator[*models.Account](ctx, finder, filters, nil)
	if err != nil {
		return err
	}
	defer func(accountGenerator *repository.ModelGenerator[*models.Account]) {
		err := accountGenerator.Close()
		if err != nil {
			log.Fatalf("failed to close account: %v", err.Error())
		}
	}(&accountGenerator)

	for accountGenerator.HasNext() {
		account, err := accountGenerator.Yield()
		if err != nil {
			log.Fatalf("failed to yield account: %v", err.Error())
		}

		for _, accountId := range account.FollowerIDs {
			followerId, err := primitive.ObjectIDFromHex(accountId)
			if err != nil {
				log.Fatalf("failed to find follower: %v", err.Error())
			}

			dataIncome := &models.Income{
				BaseModel: models.NewBaseModel(),
				Amount:    int64(defaultTotalDataIncome / account.FollowerCount),
				AccountId: followerId,
			}

			_, err = incomeRepo.Create(ctx, dataIncome)
			if err != nil {
				log.Fatalf("failed to create income: %v", err.Error())
			}

		}
	}
	return nil
}

var _ AdminServiceInterface = &AdminService{}
