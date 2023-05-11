package recipe

import "errors"

type description struct {
	value string
}

func newDescription(value string) (description, error) {
	d := description{value: value}

	if err := d.checkValidValue(); err != nil {
		return description{}, err
	}

	return d, nil
}

func (d description) checkValidValue() error {
	max := 300
	isOverUpperLimit := len(d.value) > max

	min := 0
	isOverLowerLimit := len(d.value) <= min

	if isOverUpperLimit || isOverLowerLimit {
		return errors.New("説明の値は1~300文字で入力して下さい")
	}

	return nil
}

func (d description) String() string {
	return d.value
}
