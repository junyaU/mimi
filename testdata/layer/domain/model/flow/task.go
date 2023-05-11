package flow

import "errors"

type task struct {
	value string
}

func newTask(value string) (task, error) {
	t := task{
		value: value,
	}

	if err := t.checkValidValue(); err != nil {
		return task{}, err
	}

	return t, nil
}

func (t task) checkValidValue() error {
	max := 300
	isOverUpperLimit := len(t.value) > max

	min := 0
	isOverLowerLimit := len(t.value) <= min

	if isOverUpperLimit || isOverLowerLimit {
		return errors.New("説明の値は1~300文字で入力して下さい")
	}

	return nil
}
