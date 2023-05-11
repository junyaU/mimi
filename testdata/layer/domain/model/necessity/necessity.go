package necessity

import (
	"errors"
	"github.com/junyaU/mimi/testdata/layer/domain"
	"github.com/junyaU/mimi/testdata/layer/domain/model/recipe"
)

type Necessities struct {
	Id       Id
	recipeId recipe.Id
	//ingredients []Ingredient
}

type ArgIngredients struct {
	FoodID string
	Amount int
	Unit   string
}

func RegisterNecessity(ingredients []ArgIngredients) (domain.Eventer, Id, error) {
	if len(ingredients) == 0 {
		return nil, Id{}, errors.New("ingredientが必ず１つ以上必要です")
	}

	necessityID, err := NewId()
	if err != nil {
		return nil, Id{}, err
	}

	var is []Ingredient

	for _, ingredient := range ingredients {
		i, err := newIngredient(ingredient.FoodID, ingredient.Amount, ingredient.Unit)
		if err != nil {
			return nil, Id{}, err
		}

		is = append(is, *i)
	}

	return NewRegisteredEvent(necessityID, is), necessityID, nil
}
