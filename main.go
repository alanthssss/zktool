package main

import (
	"os"
	"seed/zookeeper"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "seed"}

	zkCmd := &cobra.Command{
		Use:   "zookeeper",
		Short: "Manage Zookeeper configuration data",
	}
	zkCmd.AddCommand(zookeeper.ExportCmd)
	zkCmd.AddCommand(zookeeper.ImportCmd)
	zkCmd.AddCommand(zookeeper.UpdateCmd)

	rootCmd.AddCommand(zkCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
