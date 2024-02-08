package repository

import (
	"context"
	"errors"

	"github.com/tejiriaustin/ToW/database"
	"github.com/tejiriaustin/ToW/models"
)

type ModelGenerator[M models.Model] struct {
	ctx    context.Context
	cursor database.Cursor
}

func NewModelGenerator[M models.Model](
	ctx context.Context,
	cursorFinder FindCursor,
	filters *QueryFilter,
	sort *QuerySort,
) (ModelGenerator[M], error) {
	cur, err := cursorFinder.FindCursor(ctx, filters, sort)
	return ModelGenerator[M]{cursor: cur, ctx: ctx}, err
}

func (g *ModelGenerator[M]) HasNext() bool {
	if g.cursor == nil {
		return false
	}
	return g.cursor.Next(g.ctx)
}

func (g *ModelGenerator[M]) Yield() (M, error) {
	var m M

	if g.cursor == nil {
		return m, errors.New("cursor in generator is nil")
	}

	if err := g.cursor.Decode(&m); err != nil {
		return m, errors.New("decode failed")
	}

	return m, nil
}

func (g *ModelGenerator[M]) Close() error {
	return g.cursor.Close(g.ctx)
}
