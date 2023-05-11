package necessity

import (
	"strings"
	"testing"
)

func TestNecessityID_Identifier(t *testing.T) {
	necessityID, _ := NewId()
	id := necessityID.String()

	if !strings.Contains(id, _necessityPrefix) {
		t.Errorf("TestNecessityId_Identifier() = %v, no prefix", id)
	}
}
