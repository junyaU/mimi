package flow

import "errors"

type tips struct {
	value string
}

func newTips(value string) (tips, error) {
	t := tips{
		value: value,
	}

	if err := t.checkValidValue(); err != nil {
		return tips{}, err
	}

	return t, nil
}

func (t tips) checkValidValue() error {
	max := 100
	isOverUpperLimit := len(t.value) > max

	min := 0
	isOverLowerLimit := len(t.value) <= min

	if isOverUpperLimit || isOverLowerLimit {
		return errors.New("説明の値は1~300文字で入力して下さい")
	}

	return nil
}
