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

func (d *TableDrawer) DrawTable(rows [][]string) error {
	if len(rows) == 0 {
		return fmt.Errorf("no packages found")
	}

	if len(rows[0]) != 5 {
		return fmt.Errorf("invalid rows")
	}

	for i, row := range rows {
		nums, err := parseRowNumbers(row)
		if err != nil {
			return fmt.Errorf("error parsing numbers in row at index %d: %v", i, err)
		}

		d.table.Append(checkAndColorLimit(d, row, nums))
	}

	d.table.Render()
	return nil
}

func parseRowNumbers(row []string) ([4]int, error) {
	var nums [4]int
	for i := 1; i <= 4; i++ {
		num, err := strconv.Atoi(row[i])
		if err != nil {
			return nums, fmt.Errorf("invalid number at position %d: %v", i, err)
		}
		nums[i-1] = num
	}

	return nums, nil
}

func checkAndColorLimit(d *TableDrawer, row []string, nums [4]int) []string {
	limits := [4]int{d.maxDirectDeps, d.maxIndirectDeps, d.maxDepth, d.maxLines}

	for i, num := range nums {
		if limits[i] > 0 && num > limits[i] {
			row[0] = d.limitColor(row[0])
			row[i+1] = d.limitColor(row[i+1])
		}
	}

	return row
}
