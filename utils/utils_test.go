package utils

import (
	"testing"
)

func TestGetModuleName(t *testing.T) {
	moduleName, err := GetModuleName()
	if err != nil {
		t.Error(err)
	}

	if moduleName != "github.com/junyaU/mimi" {
		t.Error("Expected github.com/junyaU/mimi, got ", moduleName)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		slice  []string
		target string
		expect bool
	}{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{"a"}, "a", true},
		{[]string{}, "a", false},
	}

	for _, test := range tests {
		got := Contains(test.slice, test.target)
		if got != test.expect {
			t.Errorf("Contains(%v, %v) = %v, want %v", test.slice, test.target, got, test.expect)
		}

	}

}
