package inputparser

import "errors"

type Command struct {
	Path      string
	IsGraph   bool
	IsVerbose bool
}

func ParseCommand(args ...string) (Command, error) {
	if len(args) == 0 {
		return Command{}, errors.New("no arguments")
	}

	command := Command{
		Path:      args[0],
		IsVerbose: false,
		IsGraph:   false,
	}

	for _, arg := range args[1:] {
		if arg == "-g" {
			command.IsGraph = true
		}

		if arg == "-v" {
			command.IsVerbose = true
		}
	}

	return command, nil
}
