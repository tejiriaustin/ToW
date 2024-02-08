package database

import (
	"context"

	"github.com/tejiriaustin/ToW/models"
)

type InMemoryDatabase[T models.Model] struct {
	AccountsRepo []*models.Account
}

func NewDatabase[T models.Model]() Database {
	accounts := make([]*models.Account, 0)

	return &InMemoryDatabase[T]{
		AccountsRepo: accounts,
	}
}

func (i *InMemoryDatabase[T]) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) DeleteOne(ctx context.Context, filter interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) Find(ctx context.Context, filter interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) FindOne(ctx context.Context, filter interface{}) {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}) {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) InsertOne(ctx context.Context, document interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryDatabase[T]) DeleteMany(ctx context.Context, filter interface{}) error {
	//TODO implement me
	panic("implement me")
}
