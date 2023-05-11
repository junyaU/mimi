package necessity

import "testing"

func TestRegisterNecessity(t *testing.T) {
	tests := []struct {
		ingredients []ArgIngredients
		wantErr     bool
	}{
		{ingredients: []ArgIngredients{{FoodID: "dummy", Amount: 2, Unit: "kg"}}, wantErr: false},
		{ingredients: []ArgIngredients{{FoodID: "dummy", Amount: 2, Unit: "kg"}, {FoodID: "dummyId", Amount: 4, Unit: "g"}}, wantErr: false},
		{ingredients: []ArgIngredients{{FoodID: "dummy", Amount: 2, Unit: "xxx"}}, wantErr: true},
		{ingredients: []ArgIngredients{{FoodID: "dummy", Amount: 2, Unit: "ooo"}}, wantErr: true},
		{ingredients: []ArgIngredients{{FoodID: "dummy", Amount: -10, Unit: "kg"}}, wantErr: true},
		{ingredients: []ArgIngredients{{FoodID: "dummy", Amount: 0, Unit: "kg"}}, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			event, _, err := RegisterNecessity(test.ingredients)
			if err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, test.ingredients)
			}

			if event.GetEventType() != _NecessityRegisteredEvent {
				t.Fatalf("想定しているイベントタイプではありません : %#v", event.GetEventType())
			}
		} else {
			if event, _, err := RegisterNecessity(test.ingredients); err == nil {
				t.Fatalf("#%d: bad return value: %#v", i, event)
			}
		}
	}
}
