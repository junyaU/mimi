/*
Copyright © 2023 junyaU junyaadgj@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/output"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all dependencies of a package",
	Long: `Lists all the direct and indirect dependencies of a specified package.

The list provides detailed information about each dependency, 
including the number of imports and the packages that import it. 
This can be used to get a quick overview of the dependencies in 
your project. Specify the package path as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		depsChecker, err := newDepsChecker(args)
		if err != nil {
			cobra.CheckErr(err)
		}

		drawer, err := output.NewLogDrawer(depsChecker.GetNodes())
		if err != nil {
			cobra.CheckErr(fmt.Errorf("failed to create drawer: %w", err))
		}

		drawer.Draw()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
