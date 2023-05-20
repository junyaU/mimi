package cmd

import (
	"os"
	"testing"
)

func TestList(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = append(os.Args, "list", "./../testdata/layer/domain/model/creator")

	output := GetCmdOutput(t, Execute)
	want := "github.com/junyaU/mimi/testdata/layer/domain/model/creator\n  Direct Deps:\n    No direct dependency\n  Indirect Deps:\n    No indirect dependency\n"

	if output != want {
		t.Errorf("Expected output to be '%s', got '%s'", want, output)
	}
}
