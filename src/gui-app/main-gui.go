package main_gui

// TO-DO : This is the CLI version of the app. We need to create a GUI version of the app.

import (
	"fmt"
	"os"

	appdata "github.com/Orangiuss/hecate/src/core/app-data"
	mp "github.com/Orangiuss/hecate/src/core/arsenal/marketplace"
	ol "github.com/Orangiuss/hecate/src/core/one-liners"
	"github.com/spf13/cobra"
)

// Fonction principale pour le CLI
func MainGUI() {
	if err := appdata.Init(); err != nil {
		fmt.Println("Erreur d'initialisation des données:", err)
		os.Exit(1)
	}
	var oneLiners = appdata.GetOneLiners()
	// var tools = appdata.GetTools()
	var registry = appdata.GetRegistry()

	rootCmd := &cobra.Command{
		Use:   "h-cli",
		Short: "Command Line Interface for Hecate - Magic Framework",
	}

	var ShowAllOneLinersCmd = &cobra.Command{
		Use:   "ol list",
		Short: "Liste tous les one-liners disponibles",
		Run: func(cmd *cobra.Command, args []string) {
			ol.ShowOneLiners(oneLiners)
		},
	}

	var ShowAllTools = &cobra.Command{
		Use:   "tool list",
		Short: "Liste tous les outils disponibles",
		Run: func(cmd *cobra.Command, args []string) {
			mp.ShowTools(registry)
		},
	}

	// Ajouter des sous-commandes
	rootCmd.AddCommand(ShowAllOneLinersCmd)
	rootCmd.AddCommand(ShowAllTools)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
