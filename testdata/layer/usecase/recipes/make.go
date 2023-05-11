//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../_mock/usecase/$GOPACKAGE/$GOFILE
package recipes

import (
	"context"
	"errors"
	"github.com/junyaU/mimi/testdata/layer/domain/model/recipe"
	"github.com/junyaU/mimi/testdata/layer/domain/repository"
	"github.com/junyaU/mimi/testdata/layer/domain/service"
)

type MakeCommand struct {
	CreatorId   string
	Title       string
	Thumbnail   string
	Description string
	CookingTime int
	Serving     int
}

type MakeUseCase struct {
	recipeRepository repository.Recipe
	creatorService   *service.Creator
}

func NewMakeUseCase(recipeRepo repository.Recipe, creatorService *service.Creator) *MakeUseCase {
	return &MakeUseCase{
		recipeRepository: recipeRepo,
		creatorService:   creatorService,
	}
}

func (u *MakeUseCase) Exec(c MakeCommand) error {
	if !u.creatorService.Exists(c.CreatorId) {
		return errors.New("the corresponding creator does not exist")
	}

	r, err := recipe.New(c.CreatorId, c.Title, c.Thumbnail, c.Description, c.CookingTime, c.Serving)
	if err != nil {
		return err
	}

	return u.recipeRepository.Store(context.Background(), *r)
}
