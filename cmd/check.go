/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
