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

var directThreshold int
var indirectThreshold int
var depthThreshold int
var linesThreshold int

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [package path]",
	Short: "Checks the dependency thresholds of a package",
	Long: `Checks if the direct and indirect dependencies of a specified package
	exceed the provided thresholds.

	This command is useful to enforce dependency limits in your projects,
	helping to avoid overly complex package structures. Specify the package path
	as an argument, and set the thresholds using the --direct and --indirect flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkArgsNotEmpty(args); err != nil {
			cobra.CheckErr(err)
		}

		graph, err := buildDepGraph(args[0])
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := checkDepsThresholds(graph, directThreshold, indirectThreshold, depthThreshold, linesThreshold); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().IntVarP(&directThreshold, "direct", "d", 0, "Threshold for direct dependencies")
	checkCmd.Flags().IntVarP(&indirectThreshold, "indirect", "i", 0, "Threshold for indirect dependencies")
	checkCmd.Flags().IntVarP(&depthThreshold, "depth", "z", 0, "Threshold for depth of dependency graph")
	checkCmd.Flags().IntVarP(&linesThreshold, "lines", "l", 0, "Threshold for lines of code")
}

func checkDepsThresholds(graph *analysis.DepGraph, direct, indirect, depth, lines int) error {
	graph.AnalyzeIndirectDeps()

	drawer, err := output.NewLogDrawer(graph.GetNodes())
	if err != nil {
		return fmt.Errorf("failed to create drawer: %w", err)
	}

	if (direct > 0 || indirect > 0 || depth > 0 || lines > 0) && drawer.ReportExceededDeps(direct, indirect, depth, lines) {
		return fmt.Errorf("exceeded dependency threshold")
	}

	fmt.Printf("Checked dependencies successfully, no violations found.\n")
	return nil
}
