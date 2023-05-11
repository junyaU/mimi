package recipe

import (
	"errors"
)

type serving int

const (
	ONE serving = iota + 1
	//TWO
	THREE
	//FOUR
	FIVE
)

func newServing(value int) (serving, error) {
	s := serving(value)
	if err := s.checkValidNumberOfPeople(); err != nil {
		return 0, err
	}

	return s, nil
}

func (s serving) String() string {
	values := [...]string{"", "ONE", "TWO", "THREE", "FOUR", "FIVE"}
	return values[s]
}

func (s serving) checkValidNumberOfPeople() error {
	if ONE > s || FIVE < s {
		return errors.New("1~5以外の値を入れることはできません")
	}

	return nil
}
