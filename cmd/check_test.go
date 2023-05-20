package cmd

import (
	"os"
	"testing"
)

func TestCheck(t *testing.T) {
	ordArgs := os.Args
	defer func() { os.Args = ordArgs }()

	os.Args = append(os.Args, "check", "./../testdata", "-d", "5", "-i", "10")

	output := GetCmdOutput(t, Execute)

	if output != "Checked dependencies successfully, no violations found." {
		t.Errorf("Expected output to be 'Checked dependencies successfully, no violations found.', got '%s'", output)
	}
}
