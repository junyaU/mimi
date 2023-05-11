package recipe

import (
	"testing"
)

func TestServing_checkValidNumberOfPeople(t *testing.T) {
	tests := []struct {
		serving serving
		wantErr bool
	}{
		{serving: ONE, wantErr: false},
		{serving: THREE, wantErr: false},
		{serving: FIVE, wantErr: false},
		{serving: 0, wantErr: true},
		{serving: 15, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if err := test.serving.checkValidNumberOfPeople(); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, test.serving)
			}
		} else {
			if err := test.serving.checkValidNumberOfPeople(); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, test.serving)
			}
		}
	}
}
