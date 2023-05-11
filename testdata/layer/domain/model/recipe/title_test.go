package recipe

import (
	"strings"
	"testing"
)

func TestTitle_checkValidValue(t *testing.T) {
	overUpperTitle := strings.Repeat("a", 51)
	overLowerTitle := ""

	tests := []struct {
		titleVal string
		wantErr  bool
	}{
		{titleVal: "美味しいハンバーグ", wantErr: false},
		{titleVal: overUpperTitle, wantErr: true},
		{titleVal: overLowerTitle, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if title, err := newTitle(test.titleVal); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, title)
			}
		} else {
			if title, err := newTitle(test.titleVal); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, title)
			}
		}
	}
}
