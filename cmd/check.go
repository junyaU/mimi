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
var dependentThreshold int
var weightThreshold float32

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

		path := args[0]

		if dependentThreshold > 0 || weightThreshold > 0 {
			path = "./"
		}

		graph, err := buildDepGraph(path)
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := checkDepsThresholds(graph, args[0], directThreshold, indirectThreshold, depthThreshold, linesThreshold, dependentThreshold, weightThreshold); err != nil {
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
	checkCmd.Flags().IntVarP(&dependentThreshold, "dependent", "p", 0, "Threshold for dependent packages")
	checkCmd.Flags().Float32VarP(&weightThreshold, "weight", "w", 0, "Threshold for weight of dependency graph")

}

func checkDepsThresholds(graph *analysis.DepGraph, path string, direct, indirect, depth, lines, dependent int, weight float32) error {
	graph.AnalyzeIndirectDeps()

	if dependent > 0 || weight > 0 {
		graph.AnalyzeDependents()
		graph.AnalyzeWeights()
	}

	drawer, err := output.NewLogDrawer(graph.GetNodes())
	if err != nil {
		return fmt.Errorf("failed to create drawer: %w", err)
	}

	isSetOption := direct > 0 || indirect > 0 || depth > 0 || lines > 0 || dependent > 0 || weight > 0
	if isSetOption && drawer.ReportExceededDeps(path, direct, indirect, depth, lines, dependent, weight) {
		return fmt.Errorf("exceeded dependency threshold")
	}

	fmt.Printf("Checked dependencies successfully, no violations found.\n")
	return nil
}
