package flow

import (
	"encoding/json"

	"github.com/junyaU/mimi/testdata/layer/domain"
)

const _FlowWroteEvent = "FLOW_WROTE"

type WroteEvent struct {
	*domain.Event
	steps []eventStep
}

type eventStep struct {
	StepId string `json:"step_id"`
	Model  string `json:"model"`
	Task   string `json:"task"`
	Tips   string `json:"tips"`
}

func NewWroteEvent(id Id, steps []step) WroteEvent {
	var ss []eventStep

	for _, step := range steps {
		s := eventStep{
			StepId: step.stepID.identifier(),
			Model:  step.model.showModelPath(),
			Task:   step.task.value,
			Tips:   step.tips.value,
		}

		ss = append(ss, s)
	}

	return WroteEvent{
		Event: domain.NewEvent(id.String(), 0, _FlowWroteEvent),
		steps: ss,
	}
}

func (e *WroteEvent) MarshalJSON() ([]byte, error) {
	var ss []eventStep

	for _, step := range e.steps {
		var s eventStep

		s.StepId = step.StepId
		s.Model = step.Model
		s.Task = step.Task
		s.Tips = step.Tips

		ss = append(ss, s)
	}

	return json.Marshal(ss)
}

func (e *WroteEvent) UnmarshalJSON(b []byte) error {
	var u []eventStep

	if err := json.Unmarshal(b, &u); err != nil {
		return err
	}

	e.steps = u

	return nil
}
