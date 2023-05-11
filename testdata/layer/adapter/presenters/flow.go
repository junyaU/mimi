package presenters

import (
	"encoding/json"
	"github.com/junyaU/mimi/testdata/layer/usecase/flows"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Flow struct {
}

func NewFlow() *Flow {
	return &Flow{}
}

func (f *Flow) Write(c echo.Context) error {
	var req struct {
		RecipeID  string `validate:"required" form:"recipeId"`
		CreatorID string `validate:"required" form:"creatorId"`
		Steps     string `validate:"required" form:"steps"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var steps []flows.CommandSteps

	if err := json.Unmarshal([]byte(req.Steps), &steps); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//command := flows.NewWriteFlowCommand(req.RecipeID, req.CreatorID, steps)
	//if err := f.mp.Publish("flow:write", command); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}

	return c.NoContent(http.StatusOK)
}
