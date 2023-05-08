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
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)

	}

	info, err := pkginfo.New(command.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	graph := depgraph.New(info)

	if command.IsGraph {
		graphDrawer := output.NewGraphDrawer(command.MaxDirectDeps, command.MaxIndirectDeps)
		if err := graphDrawer.Draw(graph.PrintRows()); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	}

	drawer := output.NewLogDrawer(graph.GetNodes())

	if command.IsVerbose {
		drawer.Draw()
	}

	if command.IsSetMaxDeps() && drawer.ReportExceededDeps(command.MaxDirectDeps, command.MaxIndirectDeps) {
		fmt.Fprintf(os.Stderr, "Error: Exceeded max direct or indirect dependencies\n")
		os.Exit(1)
	}

	os.Exit(0)
}
