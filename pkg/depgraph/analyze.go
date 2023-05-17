package depgraph

import (
	"github.com/junyaU/mimi/pkg/pkginfo"
	"github.com/junyaU/mimi/pkg/utils"
	"strconv"
)

type Node struct {
	Package    string
	Direct     []string
	Indirect   []string
	Dependents []string
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

func (g *Graph) AnalyzeFrequencyOfUse() {
	for _, node := range g.nodes {
		for _, importedPkg := range node.Direct {
			for index := range g.nodes {
				if g.nodes[index].Package == importedPkg {
					if !utils.Contains(g.nodes[index].Dependents, node.Package) {
						g.nodes[index].Dependents = append(g.nodes[index].Dependents, node.Package)
					}
				}
			}
		}
	}
}

func (g *Graph) AnalyzeIndirectDeps() {
	for index := range g.nodes {
		visited := make(map[string]bool)
		targetIndirect := make(map[string]bool)
		findIndirectDeps(&g.nodes[index], &g.nodes[index], g.dependencyMap, targetIndirect, visited)

		for pkg := range targetIndirect {
			g.nodes[index].Indirect = append(g.nodes[index].Indirect, pkg)
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

func analyzeDirectDeps(graph *Graph, pkgs []pkginfo.Package) {
	for _, pkg := range pkgs {
		graph.nodes = append(graph.nodes, Node{
			Package: pkg.Name,
			Direct:  pkg.Imports,
		})
	}
}
