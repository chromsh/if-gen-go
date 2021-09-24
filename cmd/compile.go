package cmd

import (
	"fmt"
	"if-gen-go/compiler"
	"os"

	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "compile go file into if-gen-go",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCompile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
	//compileCmd.Flags().StringP("file", "f", "", "compile target file")
}

func runCompile(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}

	res, err := compiler.Compile(f)
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}
