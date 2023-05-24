package output

import "testing"

func TestNewGraphDrawer(t *testing.T) {
	tests := []struct {
		maxDirectDeps   int
		maxIndirectDeps int
		maxDepth        int
		maxLines        int
		wantErr         bool
	}{
		{0, 0, 0, 0, false},
		{1, -1, 1, 1, true},
		{-1, 1, 1, 1, true},
		{-1, 1, -1, 1, true},
		{5, 4, 2, 1000, false},
	}

	for _, test := range tests {
		_, err := NewTableDrawer(test.maxDirectDeps, test.maxIndirectDeps, test.maxDepth, test.maxLines)
		if err != nil && !test.wantErr {
			t.Errorf("NewGraphDrawer(%v, %v) should not return error", test.maxDirectDeps, test.maxIndirectDeps)
		}

		if err == nil && test.wantErr {
			t.Errorf("NewGraphDrawer(%v, %v) should return error", test.maxDirectDeps, test.maxIndirectDeps)
		}
	}
}

func TestGraphDrawer_Draw(t *testing.T) {
	tests := []struct {
		rows    [][]string
		wantErr bool
	}{
		{[][]string{}, true},
		{[][]string{{"a", "1"}}, true},
		{[][]string{{"a", "1", "2", "1", "2"}, {"b", "3", "4", "2", "3"}}, false},
		{[][]string{{"a", "1", "2", "1", "1"}, {"b", "z", "4", "3", "2"}}, true},
	}

	for _, test := range tests {
		graphDrawer, err := NewTableDrawer(1, 1, 1, 1)
		if err != nil {
			t.Errorf("NewGraphDrawer(1, 1, 1 ) should not return error")
		}

		err = graphDrawer.DrawTable(test.rows)
		if err != nil && !test.wantErr {
			t.Errorf("Draw(%v) should not return error", test.rows)
		}

		if err == nil && test.wantErr {
			t.Errorf("Draw(%v) should return error", test.rows)
		}
	}
}
