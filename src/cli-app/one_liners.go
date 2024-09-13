package main_cli

import (
	ol "github.com/Orangiuss/hecate/src/core/one-liners"
	"github.com/spf13/cobra"
)

func getCommandOL(oneLiners []ol.OneLiner) *cobra.Command {
	// Add the type argument for the ol.OneLiner slice

	// Define the one-liner command with oneliner or ol alias
	var oneLinerCmd = &cobra.Command{
		Use:     "oneliner",
		Short:   "Manage one-liners",
		Aliases: []string{"ol"},
	}

	var showOneLinerCmd = &cobra.Command{
		Use:   "show",
		Short: "Show a one-liner",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var oneLiner = ol.GetOneLinerByName(oneLiners, args[0])
			// Transform the one-liner into a slice of one-liners
			oneLinersss := []ol.OneLiner{oneLiner}
			ol.ShowOneLiners(oneLinersss)
		},
	}

	var listOneLinersCmd = &cobra.Command{
		Use:   "list",
		Short: "List all one-liners",
		Run: func(cmd *cobra.Command, args []string) {
			ol.ShowOneLiners(oneLiners)
		},
	}

	oneLinerCmd.AddCommand(showOneLinerCmd)
	oneLinerCmd.AddCommand(listOneLinersCmd)

	return oneLinerCmd
}

// Autres commandes similaires pour les one-liners...
