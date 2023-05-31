/*
Copyright © 2023 junyaU <junyaadgj@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/analysis"
	"github.com/junyaU/mimi/pkg/configparser"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [config file path]",
	Short: "Runs commands specified in the configuration file",
	Long: `The run command reads a configuration file and executes the commands specified in it. 
The configuration file should be in YAML format and contain a list of commands to execute.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkArgsNotEmpty(args); err != nil {
			cobra.CheckErr(err)
		}

		config, err := configparser.NewYmlConfig(args[0])
		if err != nil {
			cobra.CheckErr(err)
		}

		commands, err := config.GetCommands()
		if err != nil {
			cobra.CheckErr(err)
		}

		for i, command := range commands {
			i += 1
			fmt.Printf("command %d: %s %s \n", i, command.Name, command.Path)

			graph, err := buildDepGraph(command.Path)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("command %d failed: %w", i, err))
			}

			if err := executeCommand(command, graph); err != nil {
				cobra.CheckErr(fmt.Errorf("command %d failed: %w", i, err))
			}

			fmt.Printf("\n")
		}

		fmt.Printf("Run command completed successfully. Processed %d commands from the configuration file.\n", len(commands))
	},
}

func executeCommand(c configparser.Command, graph *analysis.DepGraph) error {
	switch c.Name {
	case "list":
		return outputDepsList(graph)
	case "table":
		return drawDepsTable(graph, c.DirectThreshold, c.IndirectThreshold, c.DepthThreshold, c.LinesThreshold)
	case "check":
		return checkDepsThresholds(
			graph,
			c.Path,
			c.DirectThreshold,
			c.IndirectThreshold,
			c.DepthThreshold,
			c.LinesThreshold,
			c.DependentThreshold,
			c.WeightThreshold,
		)
	case "deps":
		return outputDependents(graph, c.Path)
	default:
		return fmt.Errorf("invalid command name: %s", c.Name)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
