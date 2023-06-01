package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/junyaU/mimi/pkg/utils"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

const numTableIntColumns = 5

var tableHeaders = []string{"Package", "Direct Deps", "Indirect Deps", "Dependents", "Depth", "Lines", "Weight"}

type TableDrawer struct {
	table           *tablewriter.Table
	maxDirectDeps   int
	maxIndirectDeps int
	maxDepth        int
	maxLines        int
	maxDependents   int
	maxWeight       float32
	limitColor      func(a ...interface{}) string
}

func NewTableDrawer(maxDirectDeps, maxIndirectDeps, maxDependents, maxDepth, maxLines int, maxWeight float32) (*TableDrawer, error) {
	if maxDirectDeps < 0 || maxIndirectDeps < 0 || maxDepth < 0 || maxLines < 0 || maxWeight < 0 || maxDependents < 0 {
		return nil, fmt.Errorf("all limits should be positive")
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(tableHeaders)

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
		maxDependents:   maxDependents,
		maxWeight:       maxWeight,
		limitColor:      color.New(color.FgRed).SprintFunc(),
	}, nil
}

func (d *TableDrawer) DrawTable(path string, rows [][]string) error {
	if len(rows) == 0 {
		return fmt.Errorf("no packages found")
	}

	if len(rows[0]) != len(tableHeaders) {
		return fmt.Errorf("rows should have exactly %d columns", len(tableHeaders))
	}

	for i, row := range rows {
		if !utils.IsMatchedPackage(path, row[0]) {
			continue
		}

		nums, weight, err := parseRowNumbers(row)
		if err != nil {
			return fmt.Errorf("error parsing numbers in row at index %d: %v", i, err)
		}

		d.table.Append(checkAndColorLimit(d, row, nums, weight))
	}

	if d.table.NumLines() == 0 {
		return fmt.Errorf("no packages found")
	}

	d.table.Render()
	return nil
}

func parseRowNumbers(row []string) ([numTableIntColumns]int, float32, error) {
	var nums [numTableIntColumns]int

	for i := 1; i <= numTableIntColumns; i++ {
		num, err := strconv.Atoi(row[i])
		if err != nil {
			return nums, 0, fmt.Errorf("invalid number at position %d: %v", i, err)
		}
		nums[i-1] = num
	}

	parseWeight, err := strconv.ParseFloat(row[numTableIntColumns+1], 32)
	if err != nil {
		return nums, 0, fmt.Errorf("invalid weight: %v", err)
	}

	return nums, float32(parseWeight), nil
}

func checkAndColorLimit(d *TableDrawer, row []string, nums [numTableIntColumns]int, weight float32) []string {
	limits := [numTableIntColumns]int{d.maxDirectDeps, d.maxIndirectDeps, d.maxDepth, d.maxDependents, d.maxLines}

	for i, num := range nums {
		if limits[i] > 0 && num > limits[i] {
			row[0] = d.limitColor(row[0])
			row[i+1] = d.limitColor(row[i+1])
		}
	}

	if d.maxWeight > 0 && weight > d.maxWeight {
		row[0] = d.limitColor(row[0])
		row[numTableIntColumns+1] = d.limitColor(row[numTableIntColumns+1])
	}

	return row
}
