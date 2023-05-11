package necessity

import "github.com/google/uuid"

type IngredientId struct {
	prefix string
	id     string
}

const (
	ingredientPrefix = "I-i-"
)

func newIngredientId() (IngredientId, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return IngredientId{}, err
	}

	return IngredientId{
		prefix: ingredientPrefix,
		id:     id.String(),
	}, nil
}

func (i *IngredientId) identifier() string {
	return i.prefix + i.id
}
