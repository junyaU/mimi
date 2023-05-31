package output

import (
	"github.com/junyaU/mimi/pkg/analysis"
	"testing"
)

func TestNewLogDrawer(t *testing.T) {
	tests := []struct {
		nodes   []analysis.Node
		wantErr bool
	}{
		{[]analysis.Node{}, true},
		{[]analysis.Node{{"dummy", []string{}, []string{}, []string{}, 0, 10, 0.5}}, false},
	}

	for _, test := range tests {
		_, err := NewLogDrawer(test.nodes)
		if err != nil && !test.wantErr {
			t.Errorf("NewLogDrawer(%v) should not return error", test.nodes)
		}

		if err == nil && test.wantErr {
			t.Errorf("NewLogDrawer(%v) should return error", test.nodes)
		}

	}
}

func TestLogDrawer_ReportExceededDeps(t *testing.T) {
	testPackage := "github.com/junyaU/mimi/a"

	tests := []struct {
		nodes           []analysis.Node
		maxDirectDeps   int
		maxIndirectDeps int
		maxDepth        int
		maxLines        int
		maxDependents   int
		maxWeight       float32
		expect          bool
	}{
		{[]analysis.Node{
			{testPackage, []string{"a"}, []string{"b"}, []string{"c"}, 2, 10, 0.5},
		},
			1,
			1,
			3,
			100,
			10,
			0.8,
			false,
		},
		{[]analysis.Node{
			{testPackage, []string{"a", "b", "c"}, []string{"b"}, []string{"a"}, 2, 10, 0.5},
		},
			2,
			1,
			1,
			100,
			10,
			0.8,
			true,
		}, {[]analysis.Node{
			{testPackage, []string{"a", "b", "c"}, []string{"b"}, []string{"a"}, 1, 20, 0.5},
		},
			6,
			1,
			1,
			15,
			10,
			0.8,
			true,
		}, {[]analysis.Node{
			{testPackage, []string{"a", "b", "c"}, []string{"b"}, []string{"a"}, 1, 10, 0.5},
		},
			6,
			1,
			1,
			15,
			10,
			0.4,
			true,
		},
	}

	for _, test := range tests {
		logDrawer, err := NewLogDrawer(test.nodes)
		if err != nil {
			t.Errorf("NewLogDrawer(%v) should not return error", test.nodes)
		}

		fact := logDrawer.ReportExceededDeps("a", test.maxDirectDeps, test.maxIndirectDeps, test.maxDepth, test.maxLines, test.maxDependents, test.maxWeight)
		if fact != test.expect {
			t.Errorf("ReportExceededDeps(%v, %v, %v, %v, %v, %v) should return %v", test.maxDirectDeps, test.maxIndirectDeps, test.maxDepth, test.maxLines, test.maxDependents, test.maxWeight, test.expect)
		}
	}
}

func TestIsMatched(t *testing.T) {
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
		fact := isMatched(test.path, test.pkg)
		if fact != test.expect {
			t.Errorf("isMatched(%v, %v) should return %v", test.path, test.pkg, test.expect)
		}
	}
}
