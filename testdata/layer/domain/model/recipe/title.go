package recipe

import "errors"

type title struct {
	value string
}

func newTitle(value string) (title, error) {
	t := title{value: value}

	if err := t.checkValidValue(); err != nil {
		return title{}, err
	}

	return t, nil
}

func (t title) checkValidValue() error {
	max := 50
	isOverUpperLimit := len(t.value) > max

	min := 0
	isOverLowerLimit := len(t.value) <= min

	if isOverUpperLimit || isOverLowerLimit {
		return errors.New("タイトルの値は1~50文字で入力して下さい")
	}

	return nil
}

func (t title) String() string {
	return t.value
}
