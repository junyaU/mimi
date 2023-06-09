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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [package path]",
	Short: "Lists all dependencies of a package",
	Long: `Lists all the direct and indirect dependencies of a specified package.

The list provides detailed information about each dependency, 
including the number of imports and the packages that import it. 
This can be used to get a quick overview of the dependencies in 
your project. Specify the package path as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkArgsNotEmpty(args); err != nil {
			cobra.CheckErr(err)
		}

		graph, err := buildDepGraph(args[0])
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := outputDepsList(graph); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func outputDepsList(graph *analysis.DepGraph) error {
	graph.AnalyzeIndirectDeps()

	drawer, err := output.NewLogDrawer(graph.GetNodes())
	if err != nil {
		return fmt.Errorf("failed to output deps list: %w", err)
	}

	drawer.DrawList()
	return nil
}
