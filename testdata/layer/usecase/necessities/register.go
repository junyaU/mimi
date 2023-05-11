//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../_mock/usecase/$GOPACKAGE/$GOFILE
package necessities

import (
	"github.com/junyaU/mimi/testdata/layer/domain/model/necessity"
)

type RegisterHandler struct {
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{}
}

func (h *RegisterHandler) Handle(cmd RegisterNecessityCommand) error {
	_, _, err := necessity.RegisterNecessity(cmd.ArgIngredients())
	if err != nil {
		return err
	}

	return nil
}

type RegisterNecessityCommand struct {
	CreatorID   string
	RecipeID    string
	Ingredients []CommandIngredients
}

type CommandIngredients struct {
	FoodIDVal string
	AmountVal int
	UnitVal   string
}

func NewRegisterNecessityCommand(recipeID, creatorID string, ingredients []CommandIngredients) RegisterNecessityCommand {
	return RegisterNecessityCommand{
		CreatorID:   creatorID,
		RecipeID:    recipeID,
		Ingredients: ingredients,
	}
}

func (c RegisterNecessityCommand) ArgIngredients() (ingredients []necessity.ArgIngredients) {
	for _, i := range c.Ingredients {
		ingredients = append(ingredients, necessity.ArgIngredients{
			FoodID: i.FoodIDVal,
			Amount: i.AmountVal,
			Unit:   i.UnitVal,
		})
	}
	return
}
