/*
Copyright © 2023 junyaU <junyaadgj@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/configparser"
	"github.com/junyaU/mimi/pkg/depgraph"
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

			checker, err := newDepsChecker(command.Path)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("command %d failed: %w", i, err))
			}

			if err := executeCommand(command, checker); err != nil {
				cobra.CheckErr(fmt.Errorf("command %d failed: %w", i, err))
			}

			fmt.Printf("\n")
		}

		fmt.Printf("Run command completed successfully. Processed %d commands from the configuration file.\n", len(commands))
	},
}

func executeCommand(command configparser.Command, checker *depgraph.Graph) error {
	switch command.Name {
	case "list":
		return outputDepsList(checker)
	case "table":
		return drawDepsTable(checker, command.DirectThreshold, command.IndirectThreshold)
	case "check":
		return checkDepsThresholds(checker, command.DirectThreshold, command.IndirectThreshold)
	case "freq":
		return outputFrequency(checker, command.Path)
	default:
		return fmt.Errorf("invalid command name: %s", command.Name)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
