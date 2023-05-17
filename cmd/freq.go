/*
Copyright © 2023 junyaU <junyaadgj@gmail.com>
*/
package cmd

import (
	"github.com/junyaU/mimi/pkg/depgraph"
	"github.com/junyaU/mimi/pkg/output"
	"github.com/spf13/cobra"
)

// freqCmd represents the freq command
var freqCmd = &cobra.Command{
	Use:   "freq [package path]",
	Short: "Analyze the usage frequency of dependencies in a given package",
	Long: `Analyze the usage frequency of dependencies in a given package. This command inspects 
the given package and produces a breakdown of how frequently each dependency is used within it.

For example:
'freq ./mypackage' will analyze the package at path './mypackage' and output the frequency of use for each dependency.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkArgsNotEmpty(args); err != nil {
			cobra.CheckErr(err)
		}

		depsChecker, err := newDepsChecker("./")
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := outputFrequency(depsChecker, args[0]); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(freqCmd)
}

func outputFrequency(depsChecker *depgraph.Graph, path string) error {
	depsChecker.AnalyzeFrequencyOfUse()

	drawer, err := output.NewLogDrawer(depsChecker.GetNodes())
	if err != nil {
		cobra.CheckErr(err)
	}

	return drawer.DrawFrequency(path)
}
