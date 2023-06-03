package analysis

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/pkginfo"
	"github.com/junyaU/mimi/pkg/utils"
	"sort"
	"strconv"
)

const (
	directDependencyWeights   = 0.3
	indirectDependencyWeights = 0.3
	dependentWeights          = 0.2
	depthWeights              = 0.2
)

// Node represents a node in a dependency graph.
type Node struct {
	// Package is the name of the package that this node represents.
	Package string

	// Direct is a list of packages that this node directly depends on.
	Direct []string

	// Indirect is a list of packages that this node indirectly depends on.
	Indirect []string

	// Dependents is a list of packages that depend on this node.
	Dependents []string

	// Depth is the maximum depth of this node's dependency tree.
	Depth int

	// Lines is the number of lines of code in the package that this node represents.
	Lines int

	// Weight is the weighted score of this node, based on its dependencies and dependents.
	Weight float32
}

// DepGraph represents a dependency graph of a Go project.
type DepGraph struct {
	// nodes is the list of nodes in the graph.
	nodes []Node

	// dependencyMap maps package names to their respective package info.
	dependencyMap map[string]pkginfo.Package

	// directLimits, indirectLimits, dependentLimits and depthLimits are the
	// minimum and maximum number of direct dependencies, indirect dependencies,
	// dependents, and depth across all nodes in the graph.
	directLimits    *Limits
	indirectLimits  *Limits
	dependentLimits *Limits
	depthLimits     *Limits
}

type Limits struct {
	Min int
	Max int
}

func NewLimits() *Limits {
	return &Limits{
		Min: 10000,
		Max: 0,
	}
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

// PrintRows prints the rows of the dependency graph sorted based on the given sort type.
func (g *DepGraph) PrintRows(sortType SortType) [][]string {
	var rows [][]string

	var nodes []Node
	nodes = g.nodes

	switch sortType {
	case NoSort:
	case SortByWeight:
		sort.Slice(g.nodes, func(i, j int) bool {
			return g.nodes[i].Weight > g.nodes[j].Weight
		})
	}

	for _, node := range nodes {
		rows = append(rows, []string{
			node.Package,
			strconv.Itoa(len(node.Direct)),
			strconv.Itoa(len(node.Indirect)),
			strconv.Itoa(len(node.Dependents)),
			strconv.Itoa(node.Depth),
			strconv.Itoa(node.Lines),
			fmt.Sprintf("%f", node.Weight),
		})
	}
	return rows
}

// AnalyzeDependents finds and sets the dependents for each node in the dependency graph.
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

// AnalyzeIndirectDeps finds and sets the indirect dependencies for each node in the dependency graph.
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

// AnalyzePackageLines sets the number of lines of code for each node in the dependency graph.
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

// AnalyzeWeights calculates and sets the weight score for each node in the dependency graph.
func (g *DepGraph) AnalyzeWeights() {
	for index := range g.nodes {
		g.nodes[index].calculateWeightsScore(*g.directLimits, *g.indirectLimits, *g.dependentLimits, *g.depthLimits)
	}
}

// findIndirectDeps performs a depth-first search to find and calculate the maximum depth of
// indirect dependencies for a target Node in a dependency graph.
// It traverses the graph starting from the 'node' parameter, marking visited nodes to avoid
// circular dependencies. If an indirect dependency is found (a dependency that is not a direct
// dependency of the target node), it is added to 'targetIndirect'. The function recursively
// calls itself to explore the full depth of each branch in the dependency tree, returning the
// maximum depth encountered.
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

// calculateWeightsScore calculates the weight score for the node.
// The weight score is a measure of the node's importance in the dependency graph and is based on
// the number of direct dependencies, indirect dependencies, dependents, and depth of the node.
// Each of these factors is normalized with respect to the minimum and maximum values across all
// nodes in the graph, and then weighted according to predefined weights. The final weight score
// for the node is the sum of these weighted factors.
func (n *Node) calculateWeightsScore(directL, indirectL, dependentL, depthL Limits) {
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

	score := directScore + indirectScore + dependentScore + depthScore
	if score < 0 {
		score = 0
	}

	n.Weight = score
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

func (l *Limits) Update(val int) {
	if val < l.Min {
		l.Min = val
	}

	if val > l.Max {
		l.Max = val
	}
}
