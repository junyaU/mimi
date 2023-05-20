package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = append(os.Args, "run", "./../testdata/run_test")

	output := GetCmdOutput(t, Execute)
	want := "Run command completed successfully. Processed 5 commands from the configuration file."

	if !strings.Contains(output, want) {
		t.Errorf("Expected output to be '%s', got '%s'", want, output)
	}
}
