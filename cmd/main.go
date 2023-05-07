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
	command, err := inputparser.ParseCommand(os.Args[1:]...)
	if err != nil {
		panic(err)
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

	if command.IsVerbose {
		drawer := output.NewLogDrawer()
		drawer.Draw(graph.GetNodes())
	}

	os.Exit(0)
}
