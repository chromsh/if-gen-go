package cmd

import (
	"fmt"
	"if-gen-go/compiler"
	"os"

	"github.com/spf13/cobra"
)

var decompileCmd = &cobra.Command{
	Use:   "decompile",
	Short: "decompile go file into if-gen-go",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDecompile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(decompileCmd)
}

func runDecompile(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}

	res, err := compiler.Decompile(f)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
