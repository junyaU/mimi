package flow

type step struct {
	stepID stepId
	model  model
	task   task
	tips   tips
}

func newStep(modelName, taskVal, tipsVal string) (step, error) {
	model, err := newModel(modelName)
	if err != nil {
		return step{}, err
	}

	task, err := newTask(taskVal)
	if err != nil {
		return step{}, err
	}

	tips, err := newTips(tipsVal)
	if err != nil {
		return step{}, err
	}

	id, err := newStepId()
	if err != nil {
		return step{}, err
	}

	s := step{
		stepID: id,
		model:  model,
		task:   task,
		tips:   tips,
	}

	return s, nil
}

//func newStepFromEvent(s eventStep) (*step, error) {
//	model, err := newModel(s.Model)
//	if err != nil {
//		return nil, err
//	}
//
//	task, err := newTask(s.Task)
//	if err != nil {
//		return nil, err
//	}
//
//	tips, err := newTips(s.Tips)
//	if err != nil {
//		return nil, err
//	}
//
//	id, err := newStepIdFromString(s.StepId)
//	if err != nil {
//		return nil, err
//	}
//
//	return &step{
//		stepID: id,
//		model:  model,
//		task:   task,
//		tips:   tips,
//	}, nil
//}
