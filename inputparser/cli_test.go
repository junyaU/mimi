package inputparser

import "testing"

func TestCommand(t *testing.T) {
	tests := []struct {
		args    []string
		wantErr bool
	}{
		{[]string{"./path", "-g", "-v", "-mdd", "2", "-mid", "4"}, false},
		{[]string{"./path", "-g", "-v", "-mdd", "2"}, false},
		{[]string{"./path", "-g", "-v", "-mid", "4"}, false},
		{[]string{"./path", "-g", "-v"}, false},
		{[]string{"./path", "-g"}, false},
		{[]string{"./path", "-v"}, false},
		{[]string{"./path"}, false},
		{[]string{"./path", "-a"}, true},
		{[]string{"./path", "-g", "-v", "-mdd", "a"}, true},
		{[]string{"./path", "-g", "-v", "-mid", "a"}, true},
		{[]string{"./path", "-g", "-v", "-mdd"}, true},
		{[]string{"./path", "-g", "-v", "-mid"}, true},
		{[]string{"./path", "-g", "-v", "-mdd", "2", "-mid"}, true},
	}

	for _, test := range tests {
		_, err := NewCommand(test.args...)
		if err != nil && !test.wantErr {
			t.Errorf("Command(%v) should not return error", test.args)
		}

		if err == nil && test.wantErr {
			t.Errorf("Command(%v) should return error", test.args)
		}
	}
}
