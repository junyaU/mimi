package repository

import (
	"context"
	"github.com/junyaU/mimi/testdata/layer/domain/model/recipe"
)

type Recipe interface {
	Store(ctx context.Context, r recipe.Recipe) error
}
