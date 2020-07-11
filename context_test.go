package gormate_test

import (
	"context"
	"testing"

	"gorm.io/gorm"

	"github.com/ipfans/gormate"
	"github.com/stretchr/testify/assert"
)

func TestContextMate(t *testing.T) {
	ctx := context.Background()
	db, err := gormate.FromContext(ctx)
	assert.Equal(t, gormate.ErrDBNotFound, err)
	assert.Nil(t, db)

	var newDB = &gorm.DB{}
	ctx = gormate.BindContext(ctx, newDB)
	db, err = gormate.FromContext(ctx)
	assert.Nil(t, err)
	assert.Equal(t, newDB, db)
}
