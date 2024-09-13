package main_cli

import (
	"fmt"

	mp "github.com/Orangiuss/hecate/src/core/arsenal/marketplace"
	"github.com/spf13/cobra"
)

func getCommandTool(registry *mp.Registry) *cobra.Command {
	// Add the type argument for the ol.OneLiner slice

	// Define the one-liner command with oneliner or ol alias
	var oneLinerCmd = &cobra.Command{
		Use:     "tool",
		Short:   "Manage tools",
		Aliases: []string{"t, tools"},
	}

	var searchToolCmd = &cobra.Command{
		Use:   "show",
		Short: "Search a tool by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			toolName := args[0]
			tool, err := registry.FindTool(toolName)
			if err != nil {
				fmt.Printf("Error finding tool: %v\n", err)
				return
			}
			mp.ShowToolsTools(tool)
		},
	}

	var listTools = &cobra.Command{
		Use:     "list",
		Short:   "List all tools",
		Aliases: []string{"ls", "l", "all", "-l"},
		Run: func(cmd *cobra.Command, args []string) {
			mp.ShowTools(registry)
		},
	}

	oneLinerCmd.AddCommand(searchToolCmd)
	oneLinerCmd.AddCommand(listTools)

	return oneLinerCmd
}

// Autres commandes similaires pour les outils...
