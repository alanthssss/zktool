package main

import (
	"os"
	"zktool/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "zktool"}
	rootCmd.AddCommand(cmd.ExportCmd)
	rootCmd.AddCommand(cmd.ImportCmd)
	rootCmd.AddCommand(cmd.UpdateCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
