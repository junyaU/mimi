package flow

import (
	"strings"
	"testing"
)

func TestNewFlowId(t *testing.T) {
	flowID, _ := NewId()
	id := flowID.String()

	if !strings.Contains(id, _flowPrefix) {
		t.Errorf("ProcessId_identifier() = %v, no prefix", id)
	}
}

func TestNewFlowIdFromString(t *testing.T) {
	tests := []struct {
		flowID  string
		wantErr bool
	}{
		{flowID: _flowPrefix + "grj340jvrerje9", wantErr: false},
		{flowID: "yuovjckjdogrj340jvrerje9", wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if id, err := newExistingId(test.flowID); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, id)
			}
		} else {
			if id, err := newExistingId(test.flowID); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, id)
			}
		}
	}
}
