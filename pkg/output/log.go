package output

import (
	"errors"
	"github.com/fatih/color"
	"github.com/junyaU/mimi/pkg/depgraph"
)

type LogDrawer struct {
	nodes []depgraph.Node
	pkg   *color.Color
	head  *color.Color
	base  *color.Color
	fail  *color.Color
}

func NewLogDrawer(nodes []depgraph.Node) (*LogDrawer, error) {
	if len(nodes) == 0 {
		return nil, errors.New("no nodes")
	}

	return &LogDrawer{
		nodes: nodes,
		pkg:   color.New(color.FgGreen).Add(color.Bold),
		head:  color.New(color.FgCyan).Add(color.Bold),
		base:  color.New(color.FgWhite).Add(color.Bold),
		fail:  color.New(color.FgRed),
	}, nil
}

func (l *LogDrawer) Draw() {
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
	}
}

func (l *LogDrawer) ReportExceededDeps(maxDirectDeps int, maxIndirectDeps int) bool {
	exceeded := false

	for _, node := range l.nodes {
		if maxDirectDeps > 0 && len(node.Direct) > maxDirectDeps {
			l.fail.Printf("Package %s has %d direct dependencies\n", node.Package, len(node.Direct))
			exceeded = true
		}
		if maxIndirectDeps > 0 && len(node.Indirect) > maxIndirectDeps {
			l.fail.Printf("Package %s has %d indirect dependencies\n", node.Package, len(node.Indirect))
			exceeded = true
		}
	}

	return exceeded
}
