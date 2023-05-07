package output

import (
	"github.com/fatih/color"
	"github.com/junyaU/mimi/depgraph"
)

type LogDrawer struct {
	pkg  *color.Color
	head *color.Color
	base *color.Color
	fail *color.Color
}

func NewLogDrawer() *LogDrawer {
	return &LogDrawer{
		pkg:  color.New(color.FgGreen).Add(color.Bold),
		head: color.New(color.FgCyan).Add(color.Bold),
		base: color.New(color.FgWhite).Add(color.Bold),
		fail: color.New(color.FgRed),
	}
}

func (l *LogDrawer) Draw(nodes []depgraph.Node) {
	for _, node := range nodes {
		l.pkg.Println(node.Package)

		l.head.Println("  Direct Deps:")
		if len(node.To) == 0 {
			l.fail.Println("    No direct dependency")
		} else {
			for _, to := range node.To {
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
