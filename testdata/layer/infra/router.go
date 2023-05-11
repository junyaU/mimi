package infra

import (
	"github.com/go-playground/validator/v10"
	"github.com/junyaU/mimi/testdata/layer/adapter/data_handler"
	"github.com/junyaU/mimi/testdata/layer/adapter/presenters"
	"github.com/junyaU/mimi/testdata/layer/domain/service"
	"github.com/junyaU/mimi/testdata/layer/usecase/recipes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitRouter() {
	db := NewDB()

	creatorService := service.NewCreatorService(data_handler.NewUser(db))
	recipeDataHandler := data_handler.NewRecipe(db)
	recipeMake := recipes.NewMakeUseCase(recipeDataHandler, creatorService)

	recipeController := presenters.NewRecipe(recipeMake)

	e := echo.New()
	e.Validator = NewValidator()

	e.Use(middleware.Logger())

	e.POST("/recipes", recipeController.Create)

	e.Logger.Fatal(e.Start(":8081"))
}
