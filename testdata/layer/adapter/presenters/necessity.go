package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/junyaU/mimi/testdata/layer/usecase/necessities"
)

type Necessity struct {
}

func NewNecessity() *Necessity {
	return &Necessity{}
}

func (n *Necessity) Register(c echo.Context) error {
	var req struct {
		CreatorID   string `validate:"required" form:"creatorId"`
		RecipeID    string `validate:"required" form:"recipeId"`
		Ingredients string `validate:"required" form:"ingredients"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var ingredients []necessities.CommandIngredients

	if err := json.Unmarshal([]byte(req.Ingredients), &ingredients); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
