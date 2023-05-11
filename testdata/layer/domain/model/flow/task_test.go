package flow

import (
	"strings"
	"testing"
)

func TestTask_checkValidValue(t *testing.T) {
	overUpperValue := strings.Repeat("a", 301)
	overLowerValue := ""

	tests := []struct {
		task    task
		wantErr bool
	}{
		{task: task{value: "フライパンて野菜を10分炒める"}, wantErr: false},
		{task: task{value: overLowerValue}, wantErr: true},
		{task: task{value: overUpperValue}, wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if err := test.task.checkValidValue(); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, test.task)
			}
		} else {
			if err := test.task.checkValidValue(); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, test.task)
			}
		}
	}
}
