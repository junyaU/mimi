package inputparser

import (
	"errors"
	"fmt"
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

	for i := 0; i < len(args[1:]); i++ {
		arg := args[i+1]

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
				return Command{}, errors.New("mdd argument must be integer")
			}
			command.MaxDirectDeps = depsNum
			i++

		case "-mid":
			if len(args) <= i+2 {
				return Command{}, errors.New("no arguments")
			}

			depsNum, err := strconv.Atoi(args[i+2])
			if err != nil {
				return Command{}, errors.New("mid argument must be integer")
			}
			command.MaxIndirectDeps = depsNum
			i++

		default:
			return Command{}, fmt.Errorf("unknown argument: %s", arg)
		}
	}

	return command, nil
}

func (c *Command) IsSetMaxDeps() bool {
	return c.MaxDirectDeps > 0 || c.MaxIndirectDeps > 0
}
