/*
Copyright © 2023 junyaU <junyaadgj@gmail.com>
*/
package cmd

import (
	"github.com/junyaU/mimi/pkg/analysis"
	"github.com/junyaU/mimi/pkg/output"
	"github.com/spf13/cobra"
)

// depsCmd represents the freq command
var depsCmd = &cobra.Command{
	Use:   "deps [package path]",
	Short: "Displays the dependents of the specified Go package",
	Long: `The deps command of Mimi CLI is used to analyze and display the dependents of the specified Go package. 
A dependent is a package that relies on the specified package.

By identifying these dependents, you can better understand the usage and impact of a package within your Go project. This can be particularly useful for identifying potential issues or impacts before making changes to the package.

Usage example:

mimi deps ./mypackage

This command will list all packages that are dependent on the package located at ./mypackage.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkArgsNotEmpty(args); err != nil {
			cobra.CheckErr(err)
		}

		depsChecker, err := newDepsChecker("./")
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := outputDependents(depsChecker, args[0]); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(depsCmd)
}

func outputDependents(depsChecker *analysis.Graph, path string) error {
	depsChecker.AnalyzeDependents()

	drawer, err := output.NewLogDrawer(depsChecker.GetNodes())
	if err != nil {
		cobra.CheckErr(err)
	}

	return drawer.DrawDependents(path)
}
