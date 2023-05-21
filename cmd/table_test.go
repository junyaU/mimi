package cmd

import (
	"os"
	"testing"
)

func TestTable(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = append(os.Args, "table", "./../testdata/layer/domain/model/creator")

	output := GetCmdOutput(t, Execute)
	want := "+------------------------------------------------------------+-------------+---------------+-------+\n" +
		"|                          PACKAGE                           | DIRECT DEPS | INDIRECT DEPS | DEPTH |\n" +
		"+------------------------------------------------------------+-------------+---------------+-------+\n" +
		"| github.com/junyaU/mimi/testdata/layer/domain/model/creator | 0           | 0             | 0     |\n" +
		"+------------------------------------------------------------+-------------+---------------+-------+"

	if output != want {
		t.Errorf("Expected output to be '%s', got '%s'", want, output)
	}
}
