package flow

import (
	"github.com/google/uuid"
)

const _stepPrefix = "F-s-"

type stepId struct {
	prefix string
	id     string
}

func newStepId() (stepId, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return stepId{}, err
	}

	s := stepId{
		id:     id.String(),
		prefix: _stepPrefix,
	}

	return s, nil
}

//func newStepIdFromString(value string) (id stepId, err error) {
//	if !strings.Contains(value, _stepPrefix) {
//		err = errors.New("FlowIdに変換することはできません")
//		return
//	}
//
//	stepIDIndex := strings.Index(value, _stepPrefix)
//
//	id.prefix = _stepPrefix
//	id.id = value[stepIDIndex:]
//
//	return
//}

func (s stepId) identifier() string {
	return s.prefix + s.id
}
