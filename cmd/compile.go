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
	compileCmd.Flags().StringP("file", "f", "", "compile target file")
	compileCmd.PersistentFlags().String("file", "", "A help for foo")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
	fmt.Println(res)
	return nil
}
