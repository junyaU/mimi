package flow

import (
	"strings"
	"testing"
)

func TestDescription_checkValidValue(t *testing.T) {
	overUpperValue := strings.Repeat("a", 101)
	overLowerValue := ""

	tests := []struct {
		tips    tips
		wantErr bool
	}{
		{tips: tips{value: "弱めに加熱するととても美味しくなる！"}, wantErr: false},
		{tips: tips{value: overLowerValue}, wantErr: true},
		{tips: tips{value: overUpperValue}, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if err := test.tips.checkValidValue(); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, test.tips)
			}
		} else {
			if err := test.tips.checkValidValue(); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, test.tips)
			}
		}
	}
}
