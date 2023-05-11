package flow

import "testing"

func TestWriteFlow(t *testing.T) {
	tests := []struct {
		steps   []ArgSteps
		wantErr bool
	}{
		{steps: []ArgSteps{{Model: "hhhh.jpg", Task: "aaa", Tips: "sss"}}, wantErr: false},
		{steps: []ArgSteps{{Model: "ttt", Task: "aaa", Tips: "sss"}}, wantErr: true},
		{steps: []ArgSteps{{Model: "hhhh.png", Task: "", Tips: "sss"}}, wantErr: true},
		{steps: []ArgSteps{{Model: "hhhh.jpg", Task: "aaa", Tips: ""}}, wantErr: true},
		{steps: []ArgSteps{{Model: "", Task: "aaa", Tips: "sss"}}, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			event, _, err := WriteFlow(test.steps)
			if err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, test.steps)
			}

			if event.GetEventType() != _FlowWroteEvent {
				t.Fatalf("想定しているイベントタイプではありません : %#v", event.GetEventType())
			}
		} else {
			if event, _, err := WriteFlow(test.steps); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, event)
			}
		}
	}
}
