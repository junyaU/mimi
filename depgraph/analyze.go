package depgraph

import (
	"github.com/junyaU/mimi/pkginfo"
)

type Node struct {
	Package string
	From    []string
	To      []string
}

type Graph struct {
	Nodes []Node
}

func New(info []pkginfo.Info) *Graph {
	graph := &Graph{}
	for _, i := range info {
		graph.Nodes = append(graph.Nodes, Node{
			Package: i.Name,
			To:      i.Imports,
		})
	}

	for _, pkgInfo := range info {
		for _, importedPkg := range pkgInfo.Imports {
			for index := range graph.Nodes {
				if importedPkg == graph.Nodes[index].Package {
					graph.Nodes[index].From = append(graph.Nodes[index].From, pkgInfo.Name)
				}
			}
		}
	}

	return graph
}
