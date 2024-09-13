package main_cli

import (
	"fmt"
	"os"

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
			var oneLiner = ol.GetOneLinerByID(oneLiners, args[0])
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

	var searchOneLiner = &cobra.Command{
		Use:   "search",
		Short: "Search a one-liner by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			oneLinerName := args[0]
			oneLiner, _ := ol.SearchOneLiners(oneLiners, oneLinerName)
			// Transform the one-liner into a slice of one-liners
			ol.ShowOneLiners(oneLiner)
		},
	}

	var runOneLiner = &cobra.Command{
		Use:   "run",
		Short: "Run a one-liner",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			oneLinerName := args[0]
			oneLiner := ol.GetOneLinerByID(oneLiners, oneLinerName)
			ol.RunOneLiner(oneLiner, nil)
		},
	}

	var setVariables = &cobra.Command{
		Use:   "set",
		Short: "Set variables for a one-liner",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			variable := args[0]
			value := args[1]

			// Option 1: Store in environment variables
			err := os.Setenv(variable, value)
			if err != nil {
				fmt.Printf("Erreur lors de la définition de la variable d'environnement : %v\n", err)
				return
			}
			fmt.Printf("Variable %s définie à %s\n", variable, value)

			// Option 2: Store in a local configuration file or database
			// Implement your own logic to save variables persistently if needed
		},
	}

	oneLinerCmd.AddCommand(showOneLinerCmd)
	oneLinerCmd.AddCommand(listOneLinersCmd)
	oneLinerCmd.AddCommand(searchOneLiner)
	oneLinerCmd.AddCommand(runOneLiner)
	oneLinerCmd.AddCommand(setVariables)

	return oneLinerCmd
}

// Autres commandes similaires pour les one-liners...
