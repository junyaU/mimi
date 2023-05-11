//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../_mock/usecase/$GOPACKAGE/$GOFILE
package flows

import (
	"github.com/junyaU/mimi/testdata/layer/domain/model/flow"
)

type WriteHandler struct {
}

func NewWriteHandler() *WriteHandler {
	return &WriteHandler{}
}

func (h *WriteHandler) Handle(cmd WriteFlowCommand) error {
	return nil
}

type WriteFlowCommand struct {
	RecipeID  string
	CreatorID string
	Steps     []CommandSteps
}

type CommandSteps struct {
	ModelPath string
	TaskVal   string
	TipsVal   string
}

func NewWriteFlowCommand(recipeID, creatorID string, steps []CommandSteps) WriteFlowCommand {
	return WriteFlowCommand{
		RecipeID:  recipeID,
		CreatorID: creatorID,
		Steps:     steps,
	}
}

func (c WriteFlowCommand) ArgSteps() (steps []flow.ArgSteps) {
	for _, s := range c.Steps {
		steps = append(steps, flow.ArgSteps{
			Model: s.ModelPath,
			Task:  s.TaskVal,
			Tips:  s.TipsVal,
		})
	}
	return
}
