package cmd

import (
	"bytes"
	"os"
	"testing"
)

func GetCmdOutput(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() failed: %v", err)
	}

	os.Stdout = w

	fn()
	w.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("buf.ReadFrom(r) failed: %v", err)
	}

	str := buf.String()

	if str[len(str)-1] == '\n' {
		return str[:len(str)-1]
	}

	return str
}
