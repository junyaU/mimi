/*
Copyright © 2023 junyaU <junya@adgj@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/junyaU/mimi/pkg/analysis"
	"github.com/junyaU/mimi/pkg/pkginfo"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mimi",
	Short: "A CLI tool for analyzing and quantifying dependencies in Go projects.",
	Long: `mimi is a CLI tool that helps developers understand the dependencies in their Go projects. 
It analyzes the imports within the project, and provides a quantified view of both direct and indirect dependencies. 

Example usage:
	$ mimi list ./path/to/your/project
This will output a list of all packages in your project and the number of their direct and indirect dependencies.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkArgsNotEmpty(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no package path provided")
	}

	return nil
}

func newDepsChecker(path string) (*analysis.Graph, error) {
	info, err := pkginfo.New(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get package info: %w", err)
	}

	return analysis.New(info), nil
}
