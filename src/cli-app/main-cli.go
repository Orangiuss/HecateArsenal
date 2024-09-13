package main_cli

import (
	"fmt"
	"os"
	"time"

	"math/rand"

	appdata "github.com/Orangiuss/hecate/src/core/app-data"
	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

// printAscii affiche un texte en ASCII avec un style aléatoire.
func printHecate(text string) {
	// Définir les polices disponibles
	fonts := []string{
		"standard", "big", "block", "bubble", "digital",
		"fancy", "doh", "stop", "starwars",
	}

	// Initialiser le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())
	randomFont := fonts[rand.Intn(len(fonts))]

	ascii := figlet4go.NewAsciiRender()

	// Ajouter des options avec une police aléatoire
	options := figlet4go.NewRenderOptions()
	options.FontName = randomFont
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorRed, // Couleur par défaut
	}

	// Rendu du texte
	renderStr, err := ascii.RenderOpts(text, options)
	if err != nil {
		fmt.Println("Erreur lors du rendu du texte :", err)
		return
	}

	fmt.Println(renderStr)

}

// Fonction principale pour le CLI
func MainCLI() {
	if err := appdata.Init(); err != nil {
		fmt.Println("Erreur d'initialisation des données:", err)
		os.Exit(1)
	}
	var oneLiners = appdata.GetOneLiners()
	// var tools = appdata.GetTools()
	var registry = appdata.GetRegistry()

	// Afficher le logo Hecate
	printHecate("Hecate")
	fmt.Println("------------------------------------------")
	fmt.Println("CLI - Command Line Interface for Hecate")
	fmt.Println("Version:", appdata.Version)
	fmt.Println("Auteur:", appdata.Author)
	fmt.Println("------------------------------------------")

	rootCmd := &cobra.Command{
		Use:     "hecate-cli",
		Short:   "Command Line Interface for Hecate - Magic Framework",
		Aliases: []string{"h-cli", "hecate", "hct", "./hecate-cli", "./h-cli", "./hecate"},
	}

	// Add subcommands
	rootCmd.AddCommand(getCommandTool(registry))
	rootCmd.AddCommand(getCommandOL(oneLiners))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
