//go:build unit_test

package dao

import (
	"testing"

	"github.com/reymom/go-calendar-tutorial/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestEmptyPool(t *testing.T) {
	a := assert.New(t)
	pool, e := initializeEmptyPool(model.ErrReadConnectionStringEmpty)
	a.NoError(e)
	a.NotNil(pool)
}
