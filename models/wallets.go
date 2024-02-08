package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Wallet struct {
		BaseModel `json:",inline"`
	}
)

func (c Wallet) NewID() {
	//TODO implement me
	panic("implement me")
}

func (c Wallet) GetID() primitive.ObjectID {
	return c.ID
}
