package recipe

import "testing"

func TestCookingTime_checkValidTime(t *testing.T) {
	tests := []struct {
		cookingTime cookingTime
		wantErr     bool
	}{
		{cookingTime: fiveM, wantErr: false},
		{cookingTime: fiveH, wantErr: false},
		{cookingTime: twentyM, wantErr: false},
		{cookingTime: 0, wantErr: true},
		{cookingTime: 25, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if err := test.cookingTime.checkValidTime(); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, test.cookingTime)
			}
		} else {
			if err := test.cookingTime.checkValidTime(); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, test.cookingTime)
			}
		}
	}
}
