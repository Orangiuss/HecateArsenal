package main

import (
	"fmt"
	"log"
	"os"

	mp "github.com/Orangiuss/hecate/src/core/arsenal/marketplace"
	one_liners "github.com/Orangiuss/hecate/src/core/one-liners"
	"github.com/Orangiuss/hecate/src/core/utils"
)

func main() {
	if !utils.IsRoot() {
		fmt.Println("This program must be run as root.")
		os.Exit(1)
	}
	utils.PrintAscii("32", false)

	oneLiners, err := one_liners.LoadOneLiners("one-liners-templates")
	if err != nil {
		fmt.Printf("Erreur de chargement des one-liners: %v\n", err)
		return
	}
	// Afficher le nombre de one-liners chargés
	fmt.Printf("Nombre de one-liners chargés: %d\n", len(oneLiners))

	// Afficher tous les one-liners disponibles
	one_liners.ShowOneLiners(oneLiners)

	one_liners.TestOneLiner()

	utils.ShowsAllsHashingEncoding()

	// Charger le registre des outils
	registry, err := mp.LoadRegistry("src/core/arsenal/tools.json")
	if err != nil {
		log.Fatal(err)
	}

	// Afficher les outils disponibles
	mp.ShowTools(registry)

	// Afficher les outils par catégorie
	mp.ShowToolsByCategory(registry, "Reconnaissance")

	// Afficher les outils par tag
	mp.ShowToolsByTag(registry, "reconnaissance")

	// Rechercher un outil (par exemple, nmap)
	tool, err := registry.FindTool("nmap")
	if err != nil {
		log.Fatal(err)
	}

	mp.ShowTool(tool[0])

	tool2, _ := registry.FindToolByID("1")

	// Choisir le package manager (par exemple, apt)
	packageManager := "yum"

	// Obtenir la commande d'installation pour le gestionnaire de paquets choisi
	installCmd, err := tool2.GetActionByPackageManager(packageManager, "install")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Command to install %s using %s: %s\n", tool2.Name, packageManager, installCmd)

	// Exécuter la commande d'installation
	err = mp.InstallTool(*tool2, packageManager)
	if err != nil {
		log.Fatal(err)
	}

}
