package inputparser

import (
	"errors"
	"strconv"
)

type Command struct {
	Path            string
	IsGraph         bool
	IsVerbose       bool
	MaxDirectDeps   int
	MaxIndirectDeps int
}

func NewCommand(args ...string) (Command, error) {
	if len(args) == 0 {
		return Command{}, errors.New("no arguments")
	}

	command := Command{
		Path:      args[0],
		IsVerbose: false,
		IsGraph:   false,
	}

	for i, arg := range args[1:] {
		switch arg {
		case "-g":
			command.IsGraph = true

		case "-v":
			command.IsVerbose = true

		case "-mdd":
			if len(args) <= i+2 {
				return Command{}, errors.New("no arguments")
			}

			depsNum, err := strconv.Atoi(args[i+2])
			if err != nil {
				return Command{}, err
			}
			command.MaxDirectDeps = depsNum

		case "-mid":
			if len(args) <= i+2 {
				return Command{}, errors.New("no arguments")
			}

			depsNum, err := strconv.Atoi(args[i+2])
			if err != nil {
				return Command{}, err
			}
			command.MaxIndirectDeps = depsNum
		default:
		}
	}

	return command, nil
}
