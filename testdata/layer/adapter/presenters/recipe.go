package presenters

import (
	"net/http"

	"github.com/junyaU/mimi/testdata/layer/usecase/recipes"
	"github.com/labstack/echo/v4"
)

type Recipe struct {
	make *recipes.MakeUseCase
}

func NewRecipe(make *recipes.MakeUseCase) *Recipe {
	return &Recipe{
		make: make,
	}
}

func (r *Recipe) Create(c echo.Context) error {
	var req struct {
		Title       string `validate:"required,min=1,max=50" form:"title"`
		Thumbnail   string `validate:"required,max=300" form:"thumbnail"`
		Description string `validate:"required,min=1,max=300" form:"description"`
		CreatorID   string `validate:"required" form:"creatorId"`
		CookingTime int    `validate:"required,gte=1,lte=11" form:"cookingTime"`
		Serving     int    `validate:"required,gte=1,lte=5" form:"serving"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := r.make.Exec(recipes.MakeCommand{
		Title:       req.Title,
		Thumbnail:   req.Thumbnail,
		Description: req.Description,
		CreatorId:   req.CreatorID,
		CookingTime: req.CookingTime,
		Serving:     req.Serving,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (r *Recipe) ApplyForRelease(c echo.Context) error {
	var req struct {
		RecipeID  string `validate:"required" form:"recipeId"`
		CreatorID string `validate:"required" form:"creatorId"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//command := recipes.NewApplyForReleaseCommand(req.RecipeID, req.CreatorID)
	//if err := r.mp.Publish("recipe:applyForRelease", command); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}

	return c.NoContent(http.StatusOK)
}

// Publish TODO: そもそもmqから受け取って実行されるものなので移動させる
func (r *Recipe) Publish(c echo.Context) error {
	var req struct {
		RecipeID  string `validate:"required" form:"recipeId"`
		CreatorID string `validate:"required" form:"creatorId"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//command := recipes.NewPublishRecipeCommand(req.CreatorID, req.RecipeID, true)
	//if err := r.mp.Publish("recipe:publish", command); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}

	return c.NoContent(http.StatusOK)
}
