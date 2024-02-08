package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tejiriaustin/ToW/database"
	"github.com/tejiriaustin/ToW/models"
)

type (
	Creator[T models.Model] interface {
		Create(ctx context.Context, data T) (T, error)
	}
	Finder[T models.Model] interface {
		FindCursor
		FindOne(ctx context.Context, queryFilter *QueryFilter, projection *QueryProjection, findOneOptions ...*options.FindOneOptions) (T, error)
		Paginate(ctx context.Context, filters *QueryFilter, page, perPage int64, projection *QueryProjection, sort *QuerySort) ([]T, *Paginator, error)
	}
	FindCursor interface {
		FindCursor(ctx context.Context, filters *QueryFilter, sort *QuerySort) (database.Cursor, error)
	}
	Updater[T models.Model] interface {
		Update(ctx context.Context, data T) (T, error)
	}
	Deleter[T models.Model] interface {
		DeleteMany(ctx context.Context, queryFilter *QueryFilter) error
	}
)

type (
	AccountsRepoInterface[T models.Account] interface {
		Creator[T]
		Finder[T]
		Updater[T]
		Deleter[T]
	}

	WalletRepoInterface[T models.Wallet] interface {
		Creator[T]
		Finder[T]
		Updater[T]
	}

	FollowersRepoInterface[T models.Follower] interface {
		Creator[T]
		Finder[T]
		Updater[T]
	}
	IncomeRepoInterface[T models.Income] interface {
		Creator[T]
		Finder[T]
		Updater[T]
	}
)
