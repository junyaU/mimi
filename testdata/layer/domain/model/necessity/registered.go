package necessity

import (
	"encoding/json"

	"github.com/junyaU/mimi/testdata/layer/domain"
)

const _NecessityRegisteredEvent = "NECESSITY_REGISTERED"

type RegisteredEvent struct {
	*domain.Event
	ingredients []eventIngredient
}

type eventIngredient struct {
	IngredientID string `json:"ingredient_id"`
	FoodID       string `json:"food_id"`
	Amount       string `json:"amount"`
}

func NewRegisteredEvent(id Id, ingredients []Ingredient) RegisteredEvent {
	var is []eventIngredient

	for _, ingredient := range ingredients {
		i := eventIngredient{
			IngredientID: ingredient.ingredientID.identifier(),
			FoodID:       ingredient.foodId,
			Amount:       ingredient.amount.string(),
		}

		is = append(is, i)
	}

	return RegisteredEvent{
		Event:       domain.NewEvent(id.String(), 0, _NecessityRegisteredEvent),
		ingredients: is,
	}
}

func (e *RegisteredEvent) MarshalJSON() ([]byte, error) {
	var is []eventIngredient

	for _, ingredient := range e.ingredients {
		var i eventIngredient

		i.IngredientID = ingredient.IngredientID
		i.FoodID = ingredient.FoodID
		i.Amount = ingredient.Amount

		is = append(is, i)
	}

	return json.Marshal(is)
}

func (e *RegisteredEvent) UnmarshalJSON(b []byte) error {
	var u []eventIngredient

	if err := json.Unmarshal(b, &u); err != nil {
		return err
	}

	e.ingredients = u

	return nil
}
