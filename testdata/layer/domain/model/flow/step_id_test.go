package flow

import (
	"strings"
	"testing"
)

func TestStepId_Identifier(t *testing.T) {
	stepID, _ := newStepId()
	id := stepID.identifier()

	if !strings.Contains(id, _stepPrefix) {
		t.Errorf("StepId_identifier() = %v, no prefix", id)
	}
}
