package analysis

import "fmt"

type SortType int

const (
	NoSort SortType = iota + 1
	SortByWeight
)

func (s SortType) String() string {
	switch s {
	case NoSort:
		return "no sort"
	case SortByWeight:
		return "sort by weight"
	default:
		return fmt.Sprintf("invalid SortType: %d", s)
	}
}
