package output

import (
	"github.com/junyaU/mimi/pkg/depgraph"
	"testing"
)

func TestNewLogDrawer(t *testing.T) {
	tests := []struct {
		nodes   []depgraph.Node
		wantErr bool
	}{
		{[]depgraph.Node{}, true},
		{[]depgraph.Node{{"dummy", []string{}, []string{}, []string{}}}, false},
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
		nodes           []depgraph.Node
		maxDirectDeps   int
		maxIndirectDeps int
		expect          bool
	}{
		{[]depgraph.Node{
			{"a", []string{"a"}, []string{"b"}, []string{"c"}},
		},
			1,
			1,
			false,
		},
		{[]depgraph.Node{
			{"a", []string{"a", "b", "c"}, []string{"b"}, []string{}},
		},
			2,
			1,
			true,
		},
	}

	for _, test := range tests {
		logDrawer, err := NewLogDrawer(test.nodes)
		if err != nil {
			t.Errorf("NewLogDrawer(%v) should not return error", test.nodes)
		}

		fact := logDrawer.ReportExceededDeps(test.maxDirectDeps, test.maxIndirectDeps)
		if fact != test.expect {
			t.Error("ReportExceededDeps() should return true")
		}
	}
}
