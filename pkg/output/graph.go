package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type TableDrawer struct {
	table           *tablewriter.Table
	maxDirectDeps   int
	maxIndirectDeps int
	maxDepth        int
	maxLines        int
	limitColor      func(a ...interface{}) string
}

func NewTableDrawer(maxDirectDeps, maxIndirectDeps, maxDepth, maxLines int) (*TableDrawer, error) {
	if maxDirectDeps < 0 || maxIndirectDeps < 0 || maxDepth < 0 || maxLines < 0 {
		return nil, fmt.Errorf("invalid maxDirectDeps, maxIndirectDeps or maxDepth")
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Package", "Direct Deps", "Indirect Deps", "Depth", "Lines"})

	table.SetRowLine(true)
	table.SetCenterSeparator("+")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	return &TableDrawer{
		table:           table,
		maxDirectDeps:   maxDirectDeps,
		maxIndirectDeps: maxIndirectDeps,
		maxDepth:        maxDepth,
		maxLines:        maxLines,
		limitColor:      color.New(color.FgRed).SprintFunc(),
	}, nil
}

func (g *TableDrawer) DrawTable(rows [][]string) error {
	if len(rows) == 0 {
		return fmt.Errorf("no packages found")
	}

	if len(rows[0]) != 5 {
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

		linesNum, err := strconv.Atoi(row[4])
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

		if g.maxLines > 0 && linesNum > g.maxLines {
			row[0] = g.limitColor(row[0])
			row[4] = g.limitColor(row[4])
		}

		g.table.Append(row)
	}

	g.table.Render()
	return nil
}
