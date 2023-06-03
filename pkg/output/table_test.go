package output

import (
	"github.com/junyaU/mimi/pkg/analysis"
	"testing"
)

func TestNewGraphDrawer(t *testing.T) {
	tests := []struct {
		maxDirectDeps   int
		maxIndirectDeps int
		maxDependents   int
		maxDepth        int
		maxLines        int
		wantErr         bool
	}{
		{0, 0, 0, 0, 0, false},
		{1, -1, 1, 1, 1, true},
		{-1, 1, 1, 1, 1, true},
		{-1, 1, -1, 1, 1, true},
		{-1, 1, -1, -1, 1, true},
		{1, 1, 1, 1, -1, true},
		{1, 1, 1, 1, 1, false},
		{5, 4, 2, 3, 1000, false},
	}

	for _, test := range tests {
		_, err := NewTableDrawer(test.maxDirectDeps, test.maxIndirectDeps, test.maxDependents, test.maxDepth, test.maxLines)
		if err != nil && !test.wantErr {
			t.Errorf("NewGraphDrawer(%v, %v) should not return error", test.maxDirectDeps, test.maxIndirectDeps)
		}

		if err == nil && test.wantErr {
			t.Errorf("NewGraphDrawer(%v, %v) should return error", test.maxDirectDeps, test.maxIndirectDeps)
		}
	}
}

func TestGraphDrawer_Draw(t *testing.T) {
	testPackage := "github.com/junyaU/mimi/a"

	tests := []struct {
		rows    [][]string
		wantErr bool
	}{
		{[][]string{}, true},
		{[][]string{{testPackage, "1"}}, true},
		{[][]string{{testPackage, "1", "2", "1", "2", "1", "3"}, {"b", "3", "4", "2", "3", "1", "3"}}, false},
		{[][]string{{testPackage, "1", "2", "1", "1", "3", "4"}, {testPackage, "z", "4", "3", "2", "1", "3"}}, true},
	}

	for _, test := range tests {
		graphDrawer, err := NewTableDrawer(1, 1, 1, 1, 1)
		if err != nil {
			t.Errorf("NewTableDrawer(%v, %v) should not return error", 1, 1)
		}

		err = graphDrawer.DrawTable("a", test.rows, analysis.NoSort)
		if err != nil && !test.wantErr {
			t.Errorf("DrawTable(%v) should not return error", test.rows)
		}

		if err == nil && test.wantErr {
			t.Errorf("DrawTable(%v) should return error", test.rows)
		}
	}
}
