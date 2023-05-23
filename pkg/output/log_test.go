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
		{[]analysis.Node{{"dummy", []string{}, []string{}, []string{}, 0, 10}}, false},
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
	tests := []struct {
		nodes           []analysis.Node
		maxDirectDeps   int
		maxIndirectDeps int
		maxDepth        int
		expect          bool
	}{
		{[]analysis.Node{
			{"a", []string{"a"}, []string{"b"}, []string{"c"}, 2, 10},
		},
			1,
			1,
			3,
			false,
		},
		{[]analysis.Node{
			{"a", []string{"a", "b", "c"}, []string{"b"}, []string{}, 2, 10},
		},
			2,
			1,
			1,
			true,
		},
	}

	for _, test := range tests {
		logDrawer, err := NewLogDrawer(test.nodes)
		if err != nil {
			t.Errorf("NewLogDrawer(%v) should not return error", test.nodes)
		}

		fact := logDrawer.ReportExceededDeps(test.maxDirectDeps, test.maxIndirectDeps, test.maxDepth)
		if fact != test.expect {
			t.Error("ReportExceededDeps() should return true")
		}
	}
}
