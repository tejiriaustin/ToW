package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Wallet struct {
		BaseModel `json:",inline"`
		Balance   int64 `json:"balance"`
	}
)

func (c Wallet) GetID() primitive.ObjectID {
	return c.ID
}
