package output

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/junyaU/mimi/pkg/analysis"
	"github.com/junyaU/mimi/pkg/utils"
	"os"
)

// LogDrawer is responsible for printing formatted output with colors.
type LogDrawer struct {
	nodes []analysis.Node
	pkg   *color.Color
	head  *color.Color
	base  *color.Color
	fail  *color.Color
}

func NewLogDrawer(nodes []analysis.Node) (*LogDrawer, error) {
	if len(nodes) == 0 {
		return nil, errors.New("no nodes")
	}

	color.Output = os.Stdout

	return &LogDrawer{
		nodes: nodes,
		pkg:   color.New(color.FgGreen).Add(color.Bold),
		head:  color.New(color.FgCyan).Add(color.Bold),
		base:  color.New(color.FgWhite).Add(color.Bold),
		fail:  color.New(color.FgRed),
	}, nil
}

func (l *LogDrawer) DrawList() {
	for _, node := range l.nodes {
		l.pkg.Printf("%s\n", node.Package)

		l.head.Println("  Direct Deps:")
		if len(node.Direct) == 0 {
			l.fail.Println("    No direct dependency")
		} else {
			for _, to := range node.Direct {
				l.base.Println("    " + to)
			}
		}

		l.head.Println("  Indirect Deps:")
		if len(node.Indirect) == 0 {
			l.fail.Println("    No indirect dependency")
		} else {
			for _, indirect := range node.Indirect {
				l.base.Println("    " + indirect)
			}
		}
		fmt.Print("\n")
	}
}

// ReportExceededDeps checks if any package exceeds the given thresholds for dependencies,
// depth, lines, dependents, and weight, and reports any violations to the console.
// It returns true if any package exceeded the thresholds.
func (l *LogDrawer) ReportExceededDeps(path string, maxDirectDeps, maxIndirectDeps, maxDepth, maxLines, maxDependent int, weight float32) bool {
	exceeded := false

	for _, node := range l.nodes {
		if !utils.IsMatchedPackage(path, node.Package) {
			continue
		}

		if maxDirectDeps > 0 && len(node.Direct) > maxDirectDeps {
			l.fail.Printf("Package %s has %d direct dependencies\n", node.Package, len(node.Direct))
			exceeded = true
		}
		if maxIndirectDeps > 0 && len(node.Indirect) > maxIndirectDeps {
			l.fail.Printf("Package %s has %d indirect dependencies\n", node.Package, len(node.Indirect))
			exceeded = true
		}

		if maxDepth > 0 && node.Depth > maxDepth {
			l.fail.Printf("Package %s has %d depth\n", node.Package, node.Depth)
			exceeded = true
		}

		if maxLines > 0 && node.Lines > maxLines {
			l.fail.Printf("Package %s has %d lines\n", node.Package, node.Lines)
			exceeded = true
		}

		if maxDependent > 0 && len(node.Dependents) > maxDependent {
			l.fail.Printf("Package %s has %d dependents\n", node.Package, len(node.Dependents))
			exceeded = true
		}

		if weight > 0 && node.Weight > weight {
			l.fail.Printf("Package %s has %f weight\n", node.Package, node.Weight)
			exceeded = true
		}
	}

	return exceeded
}

func (l *LogDrawer) DrawDependents(path string) {
	for _, node := range l.nodes {
		if !utils.IsMatchedPackage(path, node.Package) {
			continue
		}

		l.pkg.Printf("%s\n", node.Package)
		for _, dep := range node.Dependents {
			l.base.Println("  " + dep)
		}

		if len(node.Dependents) == 0 {
			l.fail.Println("  No dependents")
		}

		fmt.Print("\n")
	}
}
