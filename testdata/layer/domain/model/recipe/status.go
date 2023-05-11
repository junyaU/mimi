package recipe

type status int

const (
	PUBLIC status = iota + 1
	PRIVATE
	LIMITED_PUBLIC
)

func (s status) String() string {
	values := [...]string{"PUBLIC", "PRIVATE", "LIMITED_PUBLIC"}
	return values[s-1]
}
