package recipes

type ApplyForReleaseHandler struct {
}

func NewApplyForReleaseHandler() *ApplyForReleaseHandler {
	return &ApplyForReleaseHandler{}
}

func (h *ApplyForReleaseHandler) Handle(cmd ApplyForReleaseCommand) error {
	return nil
}

type ApplyForReleaseCommand struct {
	RecipeID  string
	CreatorID string
}

func NewApplyForReleaseCommand(recipeID, creatorID string) ApplyForReleaseCommand {
	return ApplyForReleaseCommand{
		RecipeID:  recipeID,
		CreatorID: creatorID,
	}
}
