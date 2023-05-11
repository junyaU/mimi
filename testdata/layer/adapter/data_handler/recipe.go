package data_handler

import (
	"context"
	"time"

	"github.com/junyaU/mimi/testdata/layer/adapter"
	"github.com/junyaU/mimi/testdata/layer/domain/model/recipe"
)

type Recipe struct {
	handler adapter.DataHandler
}

type recipes struct {
	Id          string
	UserId      string
	Title       string
	Thumbnail   string
	Description string
	Status      string
	CookingTime int
	PeopleNum   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewRecipe(dh adapter.DataHandler) *Recipe {
	return &Recipe{
		handler: dh,
	}
}

func (r *Recipe) Get(id string) (*Recipe, error) {
	return nil, nil
}

func (r *Recipe) Store(ctx context.Context, re recipe.Recipe) error {
	data := recipes{
		Id:          re.Identifier(),
		UserId:      re.CreatorIdentifier(),
		Title:       re.Title(),
		Thumbnail:   re.ThumbnailPath(),
		Description: re.Description(),
		Status:      re.Status(),
		CookingTime: re.CookingTime(),
		PeopleNum:   re.Serving(),
	}

	return r.handler.Create(ctx, &data)
}
