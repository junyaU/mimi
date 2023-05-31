package analysis

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/pkginfo"
	"github.com/junyaU/mimi/pkg/utils"
	"strconv"
)

const (
	directDependencyWeights   = 0.3
	indirectDependencyWeights = 0.3
	dependentWeights          = 0.2
	depthWeights              = 0.2
)

type Node struct {
	Package    string
	Direct     []string
	Indirect   []string
	Dependents []string
	Depth      int
	Lines      int
	Weight     float32
}

type Limits struct {
	Min int
	Max int
}

type DepGraph struct {
	nodes           []Node
	dependencyMap   map[string]pkginfo.Package
	directLimits    *Limits
	indirectLimits  *Limits
	dependentLimits *Limits
	depthLimits     *Limits
}

func NewDepGraph(pkgOverview *pkginfo.PackageOverview) (*DepGraph, error) {
	if len(pkgOverview.Packages) == 0 {
		return nil, fmt.Errorf("no packages found")
	}

	dependencyMap := make(map[string]pkginfo.Package)
	for _, dependency := range pkgOverview.Dependencies {
		dependencyMap[dependency.Name] = dependency
	}

	graph := &DepGraph{
		dependencyMap:   dependencyMap,
		directLimits:    NewLimits(),
		indirectLimits:  NewLimits(),
		dependentLimits: NewLimits(),
		depthLimits:     NewLimits(),
	}

	analyzeDirectDeps(graph, pkgOverview.Packages)

	return graph, nil
}

func (g *DepGraph) GetNodes() []Node {
	return g.nodes
}

func (g *DepGraph) PrintRows() [][]string {
	var rows [][]string
	for _, node := range g.nodes {
		rows = append(rows, []string{
			node.Package,
			strconv.Itoa(len(node.Direct)),
			strconv.Itoa(len(node.Indirect)),
			strconv.Itoa(node.Depth),
			strconv.Itoa(node.Lines),
		})
	}
	return rows
}

func (g *DepGraph) AnalyzeDependents() {
	for _, node := range g.nodes {
		for _, importedPkg := range node.Direct {
			for index := range g.nodes {
				if g.nodes[index].Package == importedPkg && !utils.Contains(g.nodes[index].Dependents, node.Package) {
					g.nodes[index].Dependents = append(g.nodes[index].Dependents, node.Package)
					g.dependentLimits.Update(len(g.nodes[index].Dependents))
				}
			}
		}
	}
}

func (g *DepGraph) AnalyzeIndirectDeps() {
	for index := range g.nodes {
		visited := make(map[string]bool)
		targetIndirect := make(map[string]bool)

		depth := findIndirectDeps(&g.nodes[index], &g.nodes[index], g.dependencyMap, targetIndirect, visited, 0)

		g.nodes[index].Depth = depth
		g.depthLimits.Update(depth)

		for pkg := range targetIndirect {
			g.nodes[index].Indirect = append(g.nodes[index].Indirect, pkg)
		}

		g.indirectLimits.Update(len(g.nodes[index].Indirect))
	}
}

func (g *DepGraph) AnalyzePackageLines(projectPkgs ProjectPackages) error {
	for index := range g.nodes {
		pkg, err := projectPkgs.GetPackage(g.nodes[index].Package)
		if err != nil {
			return err
		}

		lines := pkg.GetLines()

		g.nodes[index].Lines = lines
	}

	return nil
}

func (g *DepGraph) AnalyzeWeights() {
	for index := range g.nodes {
		g.nodes[index].calculateWeightsScore(*g.directLimits, *g.indirectLimits, *g.dependentLimits, *g.depthLimits)
	}
}

// Depth-First Search
func findIndirectDeps(target *Node, node *Node, dependencyMap map[string]pkginfo.Package, targetIndirect map[string]bool, visited map[string]bool, depth int) (maxDepth int) {
	maxDepth = depth

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

		currentDepth := findIndirectDeps(target, &Node{Package: deps.Name, Direct: deps.Imports}, dependencyMap, targetIndirect, visited, depth+1)
		if currentDepth > maxDepth {
			maxDepth = currentDepth
		}

		if visited[importedPkg] {
			delete(visited, importedPkg)
		}
	}

	return
}

func (n *Node) calculateWeightsScore(directL Limits, indirectL Limits, dependentL Limits, depthL Limits) {
	normalize := func(val int, limit Limits) float32 {
		if limit.Max == limit.Min {
			return 0
		}

		return (float32(val - limit.Min)) / (float32(limit.Max - limit.Min))
	}

	directScore := normalize(len(n.Direct), directL) * directDependencyWeights
	indirectScore := normalize(len(n.Indirect), indirectL) * indirectDependencyWeights
	dependentScore := normalize(len(n.Dependents), dependentL) * dependentWeights
	depthScore := normalize(n.Depth, depthL) * depthWeights

	n.Weight = directScore + indirectScore + dependentScore + depthScore
}

func analyzeDirectDeps(graph *DepGraph, pkgs []pkginfo.Package) {
	for _, pkg := range pkgs {
		graph.nodes = append(graph.nodes, Node{
			Package: pkg.Name,
			Direct:  pkg.Imports,
		})

		graph.directLimits.Update(len(pkg.Imports))
	}
}

func NewLimits() *Limits {
	return &Limits{
		Min: 10000,
		Max: 0,
	}
}

func (l *Limits) Update(val int) {
	if val < l.Min {
		l.Min = val
	}

	if val > l.Max {
		l.Max = val
	}
}
