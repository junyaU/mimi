package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type GraphDrawer struct {
	table           *tablewriter.Table
	maxDirectDeps   int
	maxIndirectDeps int
	maxDepth        int
	limitColor      func(a ...interface{}) string
}

func NewGraphDrawer(maxDirectDeps, maxIndirectDeps, maxDepth int) (*GraphDrawer, error) {
	if maxDirectDeps < 0 || maxIndirectDeps < 0 || maxDepth < 0 {
		return nil, fmt.Errorf("invalid maxDirectDeps, maxIndirectDeps or maxDepth")
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Package", "Direct Deps", "Indirect Deps", "Depth"})

	table.SetRowLine(true)
	table.SetCenterSeparator("+")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	return &GraphDrawer{
		table:           table,
		maxDirectDeps:   maxDirectDeps,
		maxIndirectDeps: maxIndirectDeps,
		maxDepth:        maxDepth,
		limitColor:      color.New(color.FgRed).SprintFunc(),
	}, nil
}

func (g *GraphDrawer) DrawTable(rows [][]string) error {
	if len(rows) == 0 {
		return fmt.Errorf("no packages found")
	}

	if len(rows[0]) != 4 {
		return fmt.Errorf("invalid rows")
	}

	for _, row := range rows {
		directDepsNum, err := strconv.Atoi(row[1])
		if err != nil {
			return err
		}

		indirectDepsNum, err := strconv.Atoi(row[2])
		if err != nil {
			return err
		}

		depthNum, err := strconv.Atoi(row[3])
		if err != nil {
			return err
		}

		if g.maxDirectDeps > 0 && directDepsNum > g.maxDirectDeps {
			row[0] = g.limitColor(row[0])
			row[1] = g.limitColor(row[1])
		}

		if g.maxIndirectDeps > 0 && indirectDepsNum > g.maxIndirectDeps {
			row[0] = g.limitColor(row[0])
			row[2] = g.limitColor(row[2])
		}

		if g.maxDepth > 0 && depthNum > g.maxDepth {
			row[0] = g.limitColor(row[0])
			row[3] = g.limitColor(row[3])
		}

		g.table.Append(row)
	}

	g.table.Render()
	return nil
}
