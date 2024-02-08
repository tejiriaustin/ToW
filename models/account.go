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
		BaseModel        `json:",inline"`
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
		FollowerIDs      []string          `json:"follower_ids" bson:"follower_ids"`
		SubscriptionInfo *SubscriptionInfo `json:"subscription_info" bson:"subscription_info"`
	}

	AccountInfo struct {
		Id         string `json:"id" bson:"id"`
		FirstName  string `json:"first_name" bson:"first_name"`
		LastName   string `json:"last_name" bson:"last_name"`
		FullName   string `json:"full_name" bson:"full_name"`
		Email      string `json:"email" bson:"email"`
		Department string `json:"department" bson:"department"`
	}

	SubscriptionInfo struct {
		Id    string `json:"id" bson:"id"`
		Price int    `json:"price" bson:"price"`
	}
)

func (c Account) NewID() {
	//TODO implement me
	panic("implement me")
}

func (c Account) GetID() primitive.ObjectID {
	return c.ID
}

func (c Account) ValidateFollowSpend(conf *env.Config) error {
	followSpend := 100 / c.FollowerCount
	if followSpend < conf.GetInt64(env.MinimumFollowSpend) {
		return errors.New("you've exceeded your  maximum follower count - please increase your subscription to follow more")
	}
	return nil
}
