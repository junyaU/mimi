package depgraph

import (
	"github.com/junyaU/mimi/pkginfo"
	"github.com/junyaU/mimi/utils"
	"strconv"
)

type Node struct {
	Package  string
	From     []string
	To       []string
	Indirect []string
}

type Graph struct {
	nodes []Node
}

func New(info []pkginfo.Info) *Graph {
	graph := &Graph{}
	analyzeDirectDeps(graph, info)
	analyzeIndirectDeps(graph)
	return graph
}

func (g *Graph) GetNodes() []Node {
	return g.nodes
}

func (g *Graph) PrintRows() [][]string {
	var rows [][]string
	for _, node := range g.nodes {
		rows = append(rows, []string{
			node.Package,
			strconv.Itoa(len(node.To)),
			strconv.Itoa(len(node.Indirect)),
			strconv.Itoa(len(node.From)),
		})
	}
	return rows
}

func analyzeDirectDeps(graph *Graph, info []pkginfo.Info) {
	for _, i := range info {
		graph.nodes = append(graph.nodes, Node{
			Package: i.Name,
			To:      i.Imports,
		})
	}

	for _, pkgInfo := range info {
		for _, importedPkg := range pkgInfo.Imports {
			for index := range graph.nodes {
				if importedPkg == graph.nodes[index].Package {
					graph.nodes[index].From = append(graph.nodes[index].From, pkgInfo.Name)
				}
			}
		}
	}
}

func analyzeIndirectDeps(graph *Graph) {
	for index := range graph.nodes {
		targetIndirect := make(map[string]bool)
		findIndirectDeps(&graph.nodes[index], &graph.nodes[index], graph, targetIndirect)

		for pkg := range targetIndirect {
			graph.nodes[index].Indirect = append(graph.nodes[index].Indirect, pkg)
		}
	}
}

func findIndirectDeps(target *Node, node *Node, graph *Graph, targetIndirect map[string]bool) {
	for _, importedPkg := range node.To {
		for index := range graph.nodes {
			if importedPkg == graph.nodes[index].Package && target.Package != graph.nodes[index].Package {
				if !targetIndirect[graph.nodes[index].Package] && !utils.Contains(target.To, graph.nodes[index].Package) {
					targetIndirect[graph.nodes[index].Package] = true
				}

				findIndirectDeps(target, &graph.nodes[index], graph, targetIndirect)
			}
		}
	}
}
