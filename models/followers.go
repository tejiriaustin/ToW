package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Follower struct {
	BaseModel      `json:",inline"`
	UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`                   // ID of the user
	FollowerUserID primitive.ObjectID `json:"follower_user_id" bson:"follower_user_id"` // ID of the follower
}

func (f Follower) NewID() {
	//TODO implement me
	panic("implement me")
}

func (f Follower) GetID() primitive.ObjectID {
	return f.ID
}
