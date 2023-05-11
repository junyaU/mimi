package necessity

import (
	"errors"
	"strconv"
)

type amount struct {
	value int
	unit  unit
}

func newAmount(value int, unitVal string) (amount, error) {
	if value <= 0 {
		return amount{}, errors.New("valueに適切な値が入っていません")
	}

	unit, err := newUnit(unitVal)
	if err != nil {
		return amount{}, err
	}

	return amount{
		value: value,
		unit:  unit,
	}, nil
}

func (a amount) string() string {
	return strconv.Itoa(a.value) + a.unit.String()
}
