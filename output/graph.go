package output

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type GraphDrawer struct {
	table *tablewriter.Table
}

func NewGraphDrawer() *GraphDrawer {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Package", "Direct Deps", "Indirect Deps", "Exports"})

	table.SetRowLine(true)
	table.SetCenterSeparator("+")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	return &GraphDrawer{
		table: table,
	}
}

func (g *GraphDrawer) Draw(table [][]string) {
	for _, v := range table {
		g.table.Append(v)
	}
	g.table.Render()
}
