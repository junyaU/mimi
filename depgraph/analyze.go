package depgraph

import (
	"github.com/junyaU/mimi/pkginfo"
	"github.com/junyaU/mimi/utils"
	"strconv"
)

type Node struct {
	Package  string
	Direct   []string
	Indirect []string
}

type Graph struct {
	nodes         []Node
	dependencyMap map[string]pkginfo.Package
}

func New(pkgOverview *pkginfo.PackageOverview) *Graph {
	dependencyMap := make(map[string]pkginfo.Package)
	for _, dependency := range pkgOverview.Dependencies {
		dependencyMap[dependency.Name] = dependency
	}

	graph := &Graph{
		dependencyMap: dependencyMap,
	}

	analyzeDirectDeps(graph, pkgOverview.Packages)
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
			strconv.Itoa(len(node.Direct)),
			strconv.Itoa(len(node.Indirect)),
		})
	}
	return rows
}

func analyzeDirectDeps(graph *Graph, pkgs []pkginfo.Package) {
	for _, pkg := range pkgs {
		graph.nodes = append(graph.nodes, Node{
			Package: pkg.Name,
			Direct:  pkg.Imports,
		})
	}
}

func analyzeIndirectDeps(graph *Graph) {
	for index := range graph.nodes {
		visited := make(map[string]bool)
		targetIndirect := make(map[string]bool)
		findIndirectDeps(&graph.nodes[index], &graph.nodes[index], graph.dependencyMap, targetIndirect, visited)

		for pkg := range targetIndirect {
			graph.nodes[index].Indirect = append(graph.nodes[index].Indirect, pkg)
		}
	}
}

func findIndirectDeps(target *Node, node *Node, dependencyMap map[string]pkginfo.Package, targetIndirect map[string]bool, visited map[string]bool) {
	for _, importedPkg := range node.Direct {
		if visited[importedPkg] {
			continue
		}

		visited[importedPkg] = true

		deps, exist := dependencyMap[importedPkg]
		if !exist || target.Package == deps.Name {
			continue
		}

		if !targetIndirect[deps.Name] && !utils.Contains(target.Direct, deps.Name) {
			targetIndirect[deps.Name] = true
		}

		findIndirectDeps(target, &Node{Package: deps.Name, Direct: deps.Imports}, dependencyMap, targetIndirect, visited)
	}
}
