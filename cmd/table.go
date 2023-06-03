/*
Copyright © 2023 junyaU <junyaadgj@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/analysis"
	"github.com/junyaU/mimi/pkg/output"
	"github.com/spf13/cobra"
)

// tableCmd represents the graph command
var tableCmd = &cobra.Command{
	Use:   "table [package path]",
	Short: "Outputs the dependency graph as a table",
	Long: `Outputs the dependency graph of a specified package as a table.

This table provides a quick overview of both direct and indirect dependencies 
of the package, which is useful for understanding the complexity and potential 
risks in the package dependency structure. Specify the package path as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkArgsNotEmpty(args); err != nil {
			cobra.CheckErr(err)
		}

		graph, err := buildDepGraph("./")
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := drawDepsTable(graph, args[0], directThreshold, indirectThreshold, dependentThreshold, depthThreshold, linesThreshold, weightFlag); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)

	tableCmd.Flags().IntVarP(&directThreshold, "direct", "d", 0, "Threshold for direct dependencies")
	tableCmd.Flags().IntVarP(&indirectThreshold, "indirect", "i", 0, "Threshold for indirect dependencies")
	tableCmd.Flags().IntVarP(&depthThreshold, "depth", "z", 0, "Threshold for depth of dependency graph")
	tableCmd.Flags().IntVarP(&linesThreshold, "lines", "l", 0, "Threshold for lines of code")
	tableCmd.Flags().IntVarP(&dependentThreshold, "dependent", "p", 0, "Threshold for dependent packages")
	tableCmd.Flags().BoolVarP(&weightFlag, "weight", "w", false, "Show weight of dependency table")
}

func drawDepsTable(graph *analysis.DepGraph, path string, direct, indirect, dependents, depth, lines int, weightFlag bool) error {
	graph.AnalyzeIndirectDeps()
	graph.AnalyzeDependents()
	graph.AnalyzeWeights()

	drawer, err := output.NewTableDrawer(direct, indirect, dependents, depth, lines)
	if err != nil {
		return fmt.Errorf("failed to create drawer: %w", err)
	}

	sortType := analysis.NoSort
	if weightFlag {
		sortType = analysis.SortByWeight
	}

	if err := drawer.DrawTable(path, graph.PrintRows(sortType), sortType); err != nil {
		return fmt.Errorf("failed to draw table: %w", err)
	}

	return nil
}
