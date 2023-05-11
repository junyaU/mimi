package flow

import (
	"errors"
	"github.com/junyaU/mimi/testdata/layer/domain"
	"github.com/junyaU/mimi/testdata/layer/domain/model/recipe"
)

type Flow struct {
	Id       Id
	recipeId recipe.Id
	//steps   []*step
}

type ArgSteps struct {
	Model string
	Task  string
	Tips  string
}

func WriteFlow(steps []ArgSteps) (domain.Eventer, Id, error) {
	if len(steps) == 0 {
		return nil, Id{}, errors.New("stepは必ず1つ以上登録する必要があります")
	}

	flowID, err := NewId()
	if err != nil {
		return nil, Id{}, err
	}

	var sts []step

	for _, step := range steps {
		s, err := newStep(step.Model, step.Task, step.Tips)
		if err != nil {
			return nil, Id{}, err
		}

		sts = append(sts, s)
	}

	event := NewWroteEvent(flowID, sts)

	return event, flowID, nil
}

func (f Flow) Identifier() string {
	return f.Id.String()
}

func (f Flow) FlowIdentifier() Id {
	return f.Id
}
