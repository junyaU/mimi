package cmd

import (
	"os"
	"testing"
)

func TestDeps(t *testing.T) {
	ordArgs := os.Args
	defer func() { os.Args = ordArgs }()

	os.Args = append(os.Args, "deps", "./")

	output := GetCmdOutput(t, Execute)

	want := "github.com/junyaU/mimi/cmd\n  No dependents\n"

	if output != want {
		t.Errorf("Expected output to be '%s', got '%s'", want, output)
	}
}
