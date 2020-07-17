package gormate

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type contextKey struct{}

var (
	ctxKey = contextKey{}

	// ErrDBNotFound warnings gorm.DB not in context.
	ErrDBNotFound = fmt.Errorf("gorm.DB not in context")
)

// FromContext returns *gorm.DB
func FromContext(ctx context.Context) (db *gorm.DB, err error) {
	v := ctx.Value(ctxKey)
	if v == nil {
		err = ErrDBNotFound
		return
	}
	db = v.(*gorm.DB)
	return
}

// BindContext binds *gorm.DB
func BindContext(ctx context.Context, db *gorm.DB) (newCtx context.Context) {
	newCtx = context.WithValue(ctx, ctxKey, db)
	return
}