package main_cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// mainCli function: Entry point of the application
func mainCli() {
	rootCmd := &cobra.Command{
		Use:   "ha-cli",
		Short: "Command Line Interface for HecateArsenal",
	}

	// Add subcommands
	rootCmd.AddCommand(useOneLinerCmd)
	rootCmd.AddCommand(listOneLinersCmd)
	rootCmd.AddCommand(searchOneLinerCmd)
	rootCmd.AddCommand(showOneLinerCmd)
	rootCmd.AddCommand(showToolCmd)
	rootCmd.AddCommand(listToolsCmd)
	rootCmd.AddCommand(searchToolCmd)
	rootCmd.AddCommand(useToolCmd)
	rootCmd.AddCommand(installToolCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
