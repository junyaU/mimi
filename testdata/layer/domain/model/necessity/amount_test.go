package necessity

import (
	"testing"
)

func TestNewAmount(t *testing.T) {
	tests := []struct {
		value   int
		unit    string
		wantErr bool
	}{
		{value: 20, unit: "kg", wantErr: false},
		{value: 10, unit: "g", wantErr: false},
		{value: 210, unit: "l", wantErr: false},
		{value: 120, unit: "ml", wantErr: false},
		{value: 1, unit: "kg", wantErr: false},
		{value: 0, unit: "kg", wantErr: true},
		{value: 10, unit: "0", wantErr: true},
		{value: 0, unit: "kgs", wantErr: true},
		{value: 0, unit: "asd", wantErr: true},
		{value: 0, unit: "aaa", wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if amount, err := newAmount(test.value, test.unit); err != nil {
				t.Fatalf("#%d: bad return value: %#v error: %#v", i, amount, err)
			}
		} else {
			if amount, err := newAmount(test.value, test.unit); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, amount)
			}
		}
	}
}
