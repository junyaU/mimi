package necessity

type Ingredient struct {
	ingredientID IngredientId
	foodId       string // TODO: 一旦stringで代用してるがvalue objectにかえる
	amount       amount
}

func newIngredient(foodId string, amountVal int, unitVal string) (*Ingredient, error) {
	id, err := newIngredientId()
	if err != nil {
		return nil, err
	}

	amount, err := newAmount(amountVal, unitVal)
	if err != nil {
		return nil, err
	}

	return &Ingredient{
		ingredientID: id,
		foodId:       foodId,
		amount:       amount,
	}, nil
}
