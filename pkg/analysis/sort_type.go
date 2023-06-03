package analysis

import "fmt"

// SortType represents a type of sorting that can be applied.
type SortType int

const (
	// NoSort represents no sorting.
	NoSort SortType = iota + 1
	// SortByWeight represents sorting by weight.
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
