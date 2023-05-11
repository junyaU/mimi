package recipe

import (
	"testing"
)

func TestNewThumbnail(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{path: "examplePath.jpg", wantErr: false},
		{path: "1234567890.jpeg", wantErr: false},
		{path: "1234567890555.png", wantErr: false},
		{path: "1234567890.gif", wantErr: true},
		{path: "", wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if thumbnail, err := newThumbnail(test.path); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, thumbnail)
			}
		} else {
			if thumbnail, err := newThumbnail(test.path); err == nil {
				t.Fatalf("#%d: want error but no error: %#v", i, thumbnail)
			}
		}
	}
}
