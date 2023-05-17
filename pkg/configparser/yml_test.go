package configparser

import (
	"testing"
)

func TestNewYmlConfig(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{"./../../testdata/invalid", true},
		{"./../../testdata/", false},
	}

	for _, test := range tests {
		_, err := NewYmlConfig(test.path)
		if err != nil && !test.wantErr {
			t.Error(err)
		}

		if err == nil && test.wantErr {
			t.Errorf("NewYmlConfig(%v) should return error", test.path)
		}
	}
}

func TestYmlConfig_GetCommands(t *testing.T) {
	config, err := NewYmlConfig("./../../testdata/")
	if err != nil {
		t.Errorf("NewYmlConfig(%v) should not return error", "./../../testdata/")
	}

	commands, err := config.GetCommands()
	if err != nil {
		t.Errorf("GetCommands() should not return error: %v", err)
	}

	if len(commands) != 5 {
		t.Errorf("GetCommands() should return 4 commands")
	}

	for _, command := range commands {
		if command.Name == "" {
			t.Errorf("GetCommands() should return command with name")
		}
	}
}

func TestConfigCommand_IsVaild(t *testing.T) {
	tests := []struct {
		command ConfigCommand
		want    bool
	}{
		{ConfigCommand{Name: "check", Parameters: CommandParams{Path: "./../testdata"}}, true},
		{ConfigCommand{Name: "list", Parameters: CommandParams{Path: "./../testdata"}}, true},
		{ConfigCommand{Name: "table", Parameters: CommandParams{Path: "./../testdata"}}, true},
		{ConfigCommand{Name: "invalid", Parameters: CommandParams{Path: "./../testdata"}}, false},
		{ConfigCommand{Name: "check", Parameters: CommandParams{Path: ""}}, false},
	}

	for _, test := range tests {
		if test.command.IsVaild() != test.want {
			t.Errorf("IsVaild() should return %v", test.want)
		}
	}
}
