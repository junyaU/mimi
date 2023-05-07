package main

import (
	"github.com/junyaU/mimi/depgraph"
	"github.com/junyaU/mimi/output"
	"github.com/junyaU/mimi/pkginfo"
)

func main() {
	info, err := pkginfo.New("./testdata")
	if err != nil {
		panic(err)
	}

	graph := depgraph.New(info)
	drawer := output.NewGraphDrawer()
	drawer.Draw(graph.Print())
}
