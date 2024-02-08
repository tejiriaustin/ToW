package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Income struct {
	BaseModel      `json:",inline"`
	Amount         int64              `json:"amount" bson:"amount"`
	AccountId      primitive.ObjectID `json:"account_id" bson:"account_id"`
	FollowerUserID primitive.ObjectID `json:"follower_user_id" bson:"follower_user_id"`
	WeekStartDate  time.Time          `json:"week_start_date" bson:"week_start_date"`
	WeekEndDate    time.Time          `json:"week_end_date" bson:"week_end_date"`
}

func (i Income) NewID() {
	//TODO implement me
	panic("implement me")
}

func (i Income) GetID() primitive.ObjectID {
	return i.ID
}
