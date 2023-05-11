package necessity

import (
	"strings"
	"testing"
)

func TestIngredientId_identifier(t *testing.T) {
	ingredientId, _ := newIngredientId()
	id := ingredientId.identifier()

	if !strings.Contains(id, ingredientPrefix) {
		t.Errorf("IngredientId_identifier() = %v, no prefix", id)
	}
}
