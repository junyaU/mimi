package necessity

import "github.com/google/uuid"

type Id struct {
	prefix string
	id     string
}

const _necessityPrefix = "N-n-"

func NewId() (Id, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Id{}, err
	}

	return Id{
		prefix: _necessityPrefix,
		id:     id.String(),
	}, nil
}

func (i Id) String() string {
	return i.prefix + i.id
}
