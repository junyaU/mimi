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

func TestIsMatchedPackage(t *testing.T) {
	testPackage := "github.com/junyaU/mimi/testdata/layer/domain/model/creator"

	tests := []struct {
		path   string
		pkg    string
		expect bool
	}{
		{"./testdata/layer", testPackage, true},
		{"./testdata/layer/domain/invalid", testPackage, false},
		{"", testPackage, false},
		{"./testdata", "", false},
	}

	for _, test := range tests {
		fact := IsMatchedPackage(test.path, test.pkg)
		if fact != test.expect {
			t.Errorf("isMatched(%v, %v) should return %v", test.path, test.pkg, test.expect)
		}
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
