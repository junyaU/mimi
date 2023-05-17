/*
Copyright © 2023 junyaU <junyaadgj@gmail.com>
*/
package cmd

import (
	"fmt"
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

		depsChecker, err := newDepsChecker(args[0])
		if err != nil {
			cobra.CheckErr(err)
		}

		depsChecker.AnalyzeIndirectDeps()

		if err := drawDepsTable(depsChecker.PrintRows(), directThreshold, indirectThreshold); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)

	tableCmd.Flags().IntVarP(&directThreshold, "direct", "d", 0, "Threshold for direct dependencies")
	tableCmd.Flags().IntVarP(&indirectThreshold, "indirect", "i", 0, "Threshold for indirect dependencies")
}

func drawDepsTable(rows [][]string, direct, indirect int) error {
	drawer, err := output.NewGraphDrawer(direct, indirect)
	if err != nil {
		return fmt.Errorf("failed to create drawer: %w", err)
	}

	if err := drawer.DrawTable(rows); err != nil {
		return fmt.Errorf("failed to draw table: %w", err)
	}

	return nil
}
