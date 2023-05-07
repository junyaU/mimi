package main

import (
	"fmt"
	"github.com/junyaU/mimi/depgraph"
	"github.com/junyaU/mimi/inputparser"
	"github.com/junyaU/mimi/output"
	"github.com/junyaU/mimi/pkginfo"
	"os"
)

func main() {
	command, err := inputparser.NewCommand(os.Args[1:]...)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	info, err := pkginfo.New(command.Path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	graph := depgraph.New(info)

	if command.IsGraph {
		graphDrawer := output.NewGraphDrawer()
		graphDrawer.Draw(graph.Print())
	}

	drawer := output.NewLogDrawer(graph.GetNodes())

	if command.IsVerbose {
		drawer.Draw()
	}

	if command.MaxDirectDeps > 0 && drawer.ReportExceededDeps(command.MaxDirectDeps, command.MaxIndirectDeps) {
		os.Exit(1)
	}

	os.Exit(0)
}
