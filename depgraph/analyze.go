package depgraph

import (
	"github.com/junyaU/mimi/pkginfo"
	"github.com/junyaU/mimi/utils"
)

type Node struct {
	Package  string
	From     []string
	To       []string
	Indirect []string
}

type Graph struct {
	Nodes []Node
}

func New(info []pkginfo.Info) (graph *Graph) {
	analyzeDirectDeps(graph, info)
	analyzeIndirectDeps(graph)
	return
}

func analyzeDirectDeps(graph *Graph, info []pkginfo.Info) {
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
}

func analyzeIndirectDeps(graph *Graph) {
	for index := range graph.Nodes {
		targetIndirect := make(map[string]bool)
		findIndirectDeps(&graph.Nodes[index], &graph.Nodes[index], graph, targetIndirect)

		for pkg := range targetIndirect {
			graph.Nodes[index].Indirect = append(graph.Nodes[index].Indirect, pkg)
		}
	}
}

func findIndirectDeps(target *Node, node *Node, graph *Graph, targetIndirect map[string]bool) {
	for _, importedPkg := range node.To {
		for index := range graph.Nodes {
			if importedPkg == graph.Nodes[index].Package && target.Package != graph.Nodes[index].Package {
				if !targetIndirect[graph.Nodes[index].Package] && !utils.Contains(target.To, graph.Nodes[index].Package) {
					targetIndirect[graph.Nodes[index].Package] = true
				}

				findIndirectDeps(target, &graph.Nodes[index], graph, targetIndirect)
			}
		}
	}
}
