package flow

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Id struct {
	prefix string
	id     string
}

const (
	_flowPrefix = "F-f-"
)

func NewId() (Id, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Id{}, err
	}

	f := Id{
		id:     id.String(),
		prefix: _flowPrefix,
	}

	return f, nil
}

func newExistingId(value string) (Id, error) {
	if !strings.Contains(value, _flowPrefix) {
		return Id{}, errors.New("FlowIdに変換することはできません")
	}

	return Id{
		prefix: _flowPrefix,
		id:     value[strings.Index(value, _flowPrefix):],
	}, nil
}

func (f Id) String() string {
	return f.prefix + f.id
}
