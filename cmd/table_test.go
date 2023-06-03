package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = append(os.Args, "table", "./")

	output := GetCmdOutput(t, Execute)

	if !strings.Contains(output, "github.com/junyaU/mimi/cmd") {
		t.Errorf("Expected output to be '%s', got '%s'", "github.com/junyaU/mimi/cmd", output)
	}
}
