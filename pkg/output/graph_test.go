package output

import "testing"

func TestNewGraphDrawer(t *testing.T) {
	tests := []struct {
		maxDirectDeps   int
		maxIndirectDeps int
		wantErr         bool
	}{
		{0, 0, false},
		{1, -1, true},
		{-1, 1, true},
		{5, 4, false},
	}

	for _, test := range tests {
		_, err := NewGraphDrawer(test.maxDirectDeps, test.maxIndirectDeps)
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
		{[][]string{{"a", "1", "2"}, {"b", "3", "4"}}, false},
		{[][]string{{"a", "1", "2"}, {"b", "z", "4"}}, true},
	}

	for _, test := range tests {
		graphDrawer, err := NewGraphDrawer(1, 1)
		if err != nil {
			t.Errorf("NewGraphDrawer(1, 1) should not return error")
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
