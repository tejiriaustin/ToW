package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/google/uuid"

	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/models"
	"github.com/tejiriaustin/ToW/repository"
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
	GetAccountInput struct {
		AccountId string
	}
	SubscribeAccountInput struct {
		AccountId string
		Amount    int64
	}
	InvestAccountInput struct {
		InvestorId     string
		InvestorWallet *models.Wallet
		Account        *models.Account
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

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) CreateAccount(
	ctx context.Context,
	input CreateAccountInput,
	accountsRepo repository.AccountsRepoInterface[*models.Account],
) (*models.Account, error) {

	account := &models.Account{
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
	accountsRepo repository.AccountsRepoInterface[*models.Account],
) (*models.Account, error) {

	if input.AccountId == "" {
		return nil, errors.New("AccountId is required")
	}

	accountId, err := primitive.ObjectIDFromHex(input.AccountId)
	if err != nil {
		return nil, err
	}

	filter := repository.NewQueryFilter().AddFilter("_id", accountId)

	matchedAccount, err := accountsRepo.FindOne(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	matchedAccount.Status = models.StatusFrozen

	return accountsRepo.Update(ctx, matchedAccount)
}

func (s *AccountService) Subscribe(
	ctx context.Context,
	input SubscribeAccountInput,
	accountsRepo repository.AccountsRepoInterface[*models.Account],
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

	return accountsRepo.Update(ctx, matchedAccount)
}

func (s *AccountService) BuyShare(
	ctx context.Context,
	input InvestAccountInput,
	accountsRepo repository.AccountsRepoInterface[*models.Account],
	walletsRepo repository.WalletRepoInterface[*models.Wallet],
	conf *env.Config,
) error {

	if err := input.Validate(); err != nil {
		return err
	}

	input.InvestorWallet.Balance = input.InvestorWallet.Balance - 1
	_, err := walletsRepo.Update(ctx, input.InvestorWallet)
	if err != nil {
		return err
	}

	input.Account.ShareHolders = append(input.Account.ShareHolders, input.InvestorId)
	input.Account.Shares = input.Account.Shares - 1

	_, err = accountsRepo.Update(ctx, input.Account)
	if err != nil {
		return err
	}

	return nil
}

func (i InvestAccountInput) Validate() error {
	if i.Account == nil {
		return errors.New("account is required")
	}

	if i.Account.Shares < 0 {
		return errors.New("seems all the shares have been bought")
	}

	if i.InvestorId == "" {
		return errors.New("investor identifier is required")
	}

	if i.InvestorWallet.Balance <= 0 {
		return errors.New("insufficient balance")
	}

	if ok := StringContains(i.InvestorId, i.Account.FollowerIDs); !ok {
		return errors.New("you need to follow this account inorder to invest")
	}
	return nil
}

func (s *AccountService) FollowAccount(
	ctx context.Context,
	input FollowAccountInput,
	accountsRepo repository.AccountsRepoInterface[*models.Account],
	followerRepo repository.FollowersRepoInterface[*models.Follower],
	config *env.Config,
) error {

	if input.AccountId == "" {
		return errors.New("account Id is required")
	}

	if input.FollowerAccountId == "" {
		return errors.New("follower Account Id is required")
	}

	followerId, err := primitive.ObjectIDFromHex(input.FollowerAccountId)
	if err != nil {
		return err
	}

	filter := repository.
		NewQueryFilter().
		AddFilter("_id", followerId)

	follower, err := accountsRepo.FindOne(ctx, filter, nil)
	if err != nil {
		return err
	}

	if err = follower.ValidateFollowSpend(config); err != nil {
		return err
	}

	if err = follower.ValidateMaxFollowerCount(); err != nil {
		return err
	}

	accountId, err := primitive.ObjectIDFromHex(input.AccountId)
	if err != nil {
		return err
	}

	followedAccount, err := accountsRepo.FindOne(ctx,
		repository.NewQueryFilter().AddFilter("_id", accountId),
		nil)
	if err != nil {
		return err
	}

	follower.FollowerIDs = append(follower.FollowerIDs, followedAccount.ID.Hex())
	_, err = accountsRepo.Update(ctx, follower)
	if err != nil {
		return err
	}

	follow := &models.Follower{
		BaseModel:      models.NewBaseModel(),
		UserID:         followedAccount.ID,
		FollowerUserID: follower.ID,
	}
	_, err = followerRepo.Create(ctx, follow)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) TradeWally(ctx context.Context,
	input TradeWallyInput,
	accountsRepo repository.AccountsRepoInterface[*models.Account],
	config *env.Config,
) error {
	return nil
}

func (s *AccountService) GetAccount(ctx context.Context,
	input GetAccountInput,
	accountsRepo repository.AccountsRepoInterface[*models.Account],
) (*models.Account, error) {
	if input.AccountId == "" {
		return nil, errors.New("account Id is required")
	}
	return accountsRepo.FindOne(ctx, repository.NewQueryFilter().AddFilter("_id", input.AccountId), nil)
}
