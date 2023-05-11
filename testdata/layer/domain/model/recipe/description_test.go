package recipe

import (
	"strings"
	"testing"
)

func TestDescription_checkValidValue(t *testing.T) {
	overUpperTitle := strings.Repeat("a", 301)
	overLowerTitle := ""

	tests := []struct {
		titleVal string
		wantErr  bool
	}{
		{titleVal: "これは誰でも簡単に作れるとても美味しい料理です", wantErr: false},
		{titleVal: overUpperTitle, wantErr: true},
		{titleVal: overLowerTitle, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if d, err := newDescription(test.titleVal); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, d)
			}
		} else {
			if d, err := newDescription(test.titleVal); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, d)
			}
		}
	}
}
