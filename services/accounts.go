package services

import (
	"context"
	"errors"
	"github.com/tejiriaustin/ToW/env"
	"time"

	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/repository"

	"github.com/google/uuid"
)

var (
	defaultReferrals    = 5
	defaultFollowIncome = 100
)

type AccountService struct{}

var _ AccountServiceInterface = &AccountService{}

type (
	CreateAccountInput struct {
		FirstName     string
		LastName      string
		Phone         string
		DOB           time.Time
		Country       string
		ZipCode       string
		Email         string
		Profession    string
		Income        string
		Company       string
		PersonalLinks []string
		Kind          models.Kind
	}
	FreezeAccountInput struct {
		AccountId string
	}

	SubscribeAccountInput struct {
		AccountId string
		Price     string
	}
	InvestAccountInput struct {
		AccountId string
	}
	FollowAccountInput struct {
		AccountId         string
		FollowerAccountId string
	}
	TradeWallyInput struct {
		Amount           int64
		RecipientDetails string
	}
)

func (s *AccountService) CreateAccount(
	ctx context.Context,
	input CreateAccountInput,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) (models.Account, error) {

	account := models.Account{
		BaseModel:     models.NewBaseModel(),
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Phone:         input.Phone,
		DOB:           input.DOB,
		Country:       models.Country(input.Country),
		ZipCode:       input.ZipCode,
		Email:         input.Email,
		Profession:    input.Profession,
		Income:        input.Income,
		Company:       input.Company,
		PersonalLinks: input.PersonalLinks,
		Referrals:     defaultReferrals,
		FollowIncome:  defaultFollowIncome,
		Status:        models.StatusActive,
		Kind:          input.Kind,
	}

	return accountsRepo.Create(ctx, account)
}

func (s *AccountService) FreezeAccount(
	ctx context.Context,
	input FreezeAccountInput,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) (*models.Account, error) {

	if input.AccountId == "" {
		return nil, errors.New("AccountId is required")
	}

	filter := repository.NewQueryFilter().AddFilter("_id", input.AccountId)

	matchedAccount, err := accountsRepo.FindOne(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	matchedAccount.Status = models.StatusFrozen

	account, err := accountsRepo.Update(ctx, matchedAccount)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (s *AccountService) Subscribe(
	ctx context.Context,
	input SubscribeAccountInput,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) (*models.Account, error) {

	if input.AccountId == "" {
		return nil, errors.New("AccountId is required")
	}

	filter := repository.NewQueryFilter().AddFilter("_id", input.AccountId)

	matchedAccount, err := accountsRepo.FindOne(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	if matchedAccount.SubscriptionInfo == nil {
		matchedAccount.SubscriptionInfo = &models.SubscriptionInfo{
			Id:    uuid.New().String(),
			Price: 10,
		}
	} else {
		matchedAccount.SubscriptionInfo = &models.SubscriptionInfo{
			Id:    uuid.New().String(),
			Price: matchedAccount.SubscriptionInfo.Price + 5,
		}
	}

	account, err := accountsRepo.Update(ctx, matchedAccount)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AccountService) Invest(
	ctx context.Context,
	input InvestAccountInput,
	accountsRepo repository.AccountsRepoInterface[models.Account],
) error {

	if input.AccountId == "" {
		return errors.New("AccountId is required")
	}

	return nil
}

func (s *AccountService) FollowAccount(
	ctx context.Context,
	input FollowAccountInput,
	accountsRepo repository.AccountsRepoInterface[models.Account],
	config *env.Config,
) error {

	if input.AccountId == "" {
		return errors.New("account Id is required")
	}

	if input.FollowerAccountId == "" {
		return errors.New("follower Account Id is required")
	}

	var (
		followedAccount, followerAccount models.Account
		err                              error
	)

	{
		filters := repository.NewQueryFilter().AddFilter("_id", input.AccountId)
		followerAccount, err = accountsRepo.FindOne(ctx, filters, nil)
		if err != nil {
			return err
		}
	}

	if err = followerAccount.ValidateFollowSpend(config); err != nil {
		return err
	}

	{
		filters := repository.NewQueryFilter().AddFilter("_id", input.AccountId)
		followedAccount, err = accountsRepo.FindOne(ctx, filters, nil)
		if err != nil {
			return err
		}
	}

	followerAccount.FollowerIDs = append(followerAccount.FollowerIDs, followedAccount.ID.Hex())

	_, err = accountsRepo.Update(ctx, followerAccount)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountService) TradeWally(ctx context.Context,
	input TradeWallyInput,
	accountsRepo repository.AccountsRepoInterface[models.Account],
	config *env.Config,
) error {
	return nil
}
