package flow

import (
	"testing"
)

func TestNewModel(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "samplePath.jpg", wantErr: false},
		{name: "sss11111111111s.jpeg", wantErr: false},
		{name: "sss11111111111s.png", wantErr: false},
		{name: "sss11111111111s.gif", wantErr: true},
		{name: "sss1", wantErr: true},
		{name: ".jpg", wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if model, err := newModel(test.name); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, model)
			}
		} else {
			if model, err := newModel(test.name); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, model)
			}
		}
	}
}
