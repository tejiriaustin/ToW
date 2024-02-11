package models

import (
	"errors"
	"github.com/tejiriaustin/ToW/env"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Country string // country code for supported countries
	Kind    string // Kind of account
	Status  string // status of an account
)

var (
	CustomerAccount Kind = "CUSTOMER"
	AdminAccount    Kind = "INTERNAL"

	StatusFrozen Status = "FROZEN"
	StatusActive Status = "ACTIVE"
)

type (
	Account struct {
		BaseModel        `bson:",inline"`
		FirstName        string            `json:"first_name" bson:"first_name"`
		LastName         string            `json:"last_name" bson:"last_name"`
		Phone            string            `json:"phone" bson:"phone"`
		DOB              time.Time         `json:"dob" bson:"dob"`
		Country          Country           `json:"country" bson:"country"`
		ZipCode          string            `json:"zip_code" bson:"zip_code"`
		Email            string            `json:"email" bson:"email"`
		Profession       string            `json:"profession" bson:"profession"`
		Income           string            `json:"income" bson:"income"`
		Company          string            `json:"company" bson:"company"`
		PersonalLinks    []string          `json:"personal_links" bson:"personal_links"`
		Referrals        int               `json:"referrals" bson:"referrals"`
		FollowIncome     int               `json:"follow_income" bson:"follow_income"`
		Token            string            `json:"-" bson:"-"`
		Kind             Kind              `json:"kind" bson:"kind"`
		Status           Status            `json:"status" bson:"status"`
		FollowerCount    int64             `json:"follower_count" bson:"follower_count"`
		Shares           int32             `json:"shares" bson:"shares"`
		MaxFollowerCount int64             `json:"max_follower_count" bson:"max_follower_count"`
		FollowerIDs      []string          `json:"follower_ids" bson:"follower_ids"`
		SubscriptionInfo *SubscriptionInfo `json:"subscription_info" bson:"subscription_info"`
		AcquiredShares   []string          `json:"acquired_shares" bson:"acquired_shares"`
		ShareHolders     []string          `json:"share_holders" json:"share_holders"`
	}

	AccountInfo struct {
		Id        string `json:"_id" bson:"_id"`
		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
		FullName  string `json:"full_name" bson:"full_name"`
		Email     string `json:"email" bson:"email"`
		Kind      string `json:"kind" bson:"kind"`
	}

	SubscriptionInfo struct {
		Id    string `json:"id" bson:"id"`
		Price int    `json:"price" bson:"price"`
	}
	Shares struct {
		AccountId string `json:"account_id" bson:"account_id"`
	}
)

func (c Account) NewID() {
	//TODO implement me
	panic("implement me")
}

func (c Account) GetID() primitive.ObjectID {
	return c.ID
}

func (c Account) FollowSpend() int64 {
	if c.FollowerCount == 0 {
		return 100
	}
	return 100 / c.FollowerCount
}

func (c Account) ValidateFollowSpend(conf *env.Config) error {
	if c.FollowSpend() < conf.GetInt64(env.MinimumFollowSpend) {
		return errors.New("following this user violates the current minimum follow spend limit - please increase your subscription to follow more")
	}
	return nil
}

func (c Account) ValidateMaxFollowerCount() error {
	if c.FollowerCount+1 >= c.MaxFollowerCount {
		return errors.New("you've exceeded your maximum follower count - please increase your subscription to follow more")
	}
	return nil
}

func (k Kind) String() string {
	return string(k)
}

func (c Account) GetAccountInfo() *AccountInfo {
	return &AccountInfo{
		Id:        c.ID.Hex(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		FullName:  c.LastName,
		Email:     c.Email,
		Kind:      c.Kind.String(),
	}
}
