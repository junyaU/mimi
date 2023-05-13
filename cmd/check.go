/*
Copyright © 2023 junyaU junyaadgj@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/depgraph"
	"github.com/junyaU/mimi/pkg/output"
	"github.com/junyaU/mimi/pkg/pkginfo"
	"github.com/spf13/cobra"
)

var directThreshold int
var indirectThreshold int

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks the dependency thresholds of a package",
	Long: `Checks if the direct and indirect dependencies of a specified package
	exceed the provided thresholds.

	This command is useful to enforce dependency limits in your projects,
	helping to avoid overly complex package structures. Specify the package path
	as an argument, and set the thresholds using the --direct and --indirect flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr("path is required")
		}

		info, err := pkginfo.New(args[0])
		if err != nil {
			cobra.CheckErr(fmt.Errorf("failed to get package info: %w", err))
		}

		graph := depgraph.New(info)

		drawer, err := output.NewLogDrawer(graph.GetNodes())
		if err != nil {
			cobra.CheckErr(fmt.Errorf("failed to create drawer: %w", err))
		}

		if (directThreshold > 0 || indirectThreshold > 0) && drawer.ReportExceededDeps(directThreshold, indirectThreshold) {
			cobra.CheckErr(fmt.Errorf("exceeded dependency threshold"))
		}

		fmt.Printf("Checked dependencies successfully, no violations found.\n")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().IntVarP(&directThreshold, "direct", "d", 0, "Threshold for direct dependencies")
	checkCmd.Flags().IntVarP(&indirectThreshold, "indirect", "i", 0, "Threshold for indirect dependencies")
}
