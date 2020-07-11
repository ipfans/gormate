package gormate

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type contextKey struct{}

var (
	ctxKey = contextKey{}

	ErrDBNotFound = fmt.Errorf("gorm.DB not in context")
)

func FromContext(ctx context.Context) (db *gorm.DB, err error) {
	v := ctx.Value(ctxKey)
	if v == nil {
		err = ErrDBNotFound
		return
	}
	db = v.(*gorm.DB)
	return
}

func BindContext(ctx context.Context, db *gorm.DB) (newCtx context.Context) {
	newCtx = context.WithValue(ctx, ctxKey, db)
	return
}