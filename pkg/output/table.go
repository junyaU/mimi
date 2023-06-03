package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/junyaU/mimi/pkg/analysis"
	"github.com/junyaU/mimi/pkg/utils"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

// The number of table columns that contain integers.
const numTableIntColumns = 5

var tableHeaders = []string{"Package", "Direct Deps", "Indirect Deps", "Dependents", "Depth", "Lines", "Weight"}

// TableDrawer is responsible for drawing tables with dependency data.
type TableDrawer struct {
	table           *tablewriter.Table
	maxDirectDeps   int
	maxIndirectDeps int
	maxDepth        int
	maxLines        int
	maxDependents   int
	lowColor        func(a ...interface{}) string
	middleColor     func(a ...interface{}) string
	highColor       func(a ...interface{}) string
}

func NewTableDrawer(maxDirectDeps, maxIndirectDeps, maxDependents, maxDepth, maxLines int) (*TableDrawer, error) {
	if maxDirectDeps < 0 || maxIndirectDeps < 0 || maxDepth < 0 || maxLines < 0 || maxDependents < 0 {
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
		lowColor:        color.New(color.FgGreen).SprintFunc(),
		middleColor:     color.New(color.FgYellow).SprintFunc(),
		highColor:       color.New(color.FgRed).SprintFunc(),
	}, nil
}

// DrawTable renders a table with given data, sorted by given sort type.
// It colors the table based on the sort type: by limit or by weight.
// It returns an error if no packages found or any row does not match the table headers.
func (d *TableDrawer) DrawTable(path string, rows [][]string, sortType analysis.SortType) error {
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

		nums, weight, err := parseRowNumbersAndWeight(row)
		if err != nil {
			return fmt.Errorf("error parsing numbers in row at index %d: %v", i, err)
		}

		var parsedRow []string
		switch sortType {
		case analysis.SortByWeight:
			parsedRow = d.colorWeight(row, weight)
		case analysis.NoSort:
			parsedRow = d.colorLimit(row, nums)
		}

		d.table.Append(parsedRow)
	}

	if d.table.NumLines() == 0 {
		return fmt.Errorf("no packages found")
	}

	d.table.Render()
	return nil
}

func (d *TableDrawer) colorLimit(row []string, nums [numTableIntColumns]int) []string {
	limits := [numTableIntColumns]int{d.maxDirectDeps, d.maxIndirectDeps, d.maxDepth, d.maxDependents, d.maxLines}

	for i, num := range nums {
		if limits[i] > 0 && num > limits[i] {
			row[0] = d.highColor(row[0])
			row[i+1] = d.highColor(row[i+1])
		}
	}

	return row
}

func (d *TableDrawer) colorWeight(row []string, weight float32) []string {
	const lowWeightLimit = 0.3
	const middleWeightLimit = 0.7
	const highWeightLimit = 1.0

	if weight < lowWeightLimit {
		row[0] = d.lowColor(row[0])
		row[numTableIntColumns+1] = d.lowColor(row[numTableIntColumns+1])
	} else if weight < middleWeightLimit {
		row[0] = d.middleColor(row[0])
		row[numTableIntColumns+1] = d.middleColor(row[numTableIntColumns+1])
	} else if weight < highWeightLimit {
		row[0] = d.highColor(row[0])
		row[numTableIntColumns+1] = d.highColor(row[numTableIntColumns+1])
	}

	return row
}

func parseRowNumbersAndWeight(row []string) ([numTableIntColumns]int, float32, error) {
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
