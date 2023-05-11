package necessity

import "errors"

type unit int

const (
	kg unit = iota + 1
	g
	l
	ml
)

func newUnit(value string) (u unit, err error) {
	switch value {
	case "kg":
		u = kg
	case "g":
		u = g
	case "l":
		u = l
	case "ml":
		u = ml
	default:
		err = errors.New("この単位は存在しません")
	}

	return
}

func (u unit) String() string {
	data := [...]string{"", "kg", "g", "l", "ml"}
	return data[u]
}
