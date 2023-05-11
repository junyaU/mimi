//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../_mock/usecase/$GOPACKAGE/$GOFILE
package recipes

type PublishHandler struct {
}

func NewPublishHandler() *PublishHandler {
	return &PublishHandler{}
}

func (h *PublishHandler) Handle(cmd PublishRecipeCommand) error {
	return nil
}

type PublishRecipeCommand struct {
	CreatorID         string
	RecipeID          string
	PostingPermission bool
}

func NewPublishRecipeCommand(creatorId, recipeID string, permission bool) PublishRecipeCommand {
	return PublishRecipeCommand{
		CreatorID:         creatorId,
		RecipeID:          recipeID,
		PostingPermission: permission,
	}
}
