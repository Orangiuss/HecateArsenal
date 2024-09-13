package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Orangiuss/hecate/src/core/arsenal/marketplace"
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

	data := "example"
	fmt.Println("MD5:", utils.HashMD5(data))
	fmt.Println("SHA256:", utils.HashSHA256(data))

	hexData := utils.ToHex(data)
	fmt.Println("To Hex:", hexData)
	decodedHex, _ := utils.FromHex(hexData)
	fmt.Println("From Hex:", decodedHex)

	hexDump := utils.ToHexDump(data)
	fmt.Println("To Hex Dump:", hexDump)
	decodedHexDump, _ := utils.FromHexDump(hexDump)
	fmt.Println("From Hex Dump:", decodedHexDump)

	match, _ := utils.RegexMatch(`[a-z]+`, data)
	fmt.Println("Regex Match:", match)

	replaced, _ := utils.RegexReplace(`e`, data, "E")
	fmt.Println("Regex Replace:", replaced)

	// Charger le registre des outils
	registry, err := marketplace.LoadRegistry("src/core/arsenal/tools.json")
	if err != nil {
		log.Fatal(err)
	}

	// Afficher les outils disponibles
	marketplace.ShowTools(registry.Tools)

	// Rechercher un outil (par exemple, nmap)
	tool, err := registry.FindTool("nmap")
	if err != nil {
		log.Fatal(err)
	}

	// Choisir le package manager (par exemple, apt)
	packageManager := "yum"

	// Obtenir la commande d'installation pour le gestionnaire de paquets choisi
	installCmd, err := tool.GetActionByPackageManager(packageManager, "install")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Command to install %s using %s: %s\n", tool.Name, packageManager, installCmd)

	// Exécuter la commande d'installation
	err = marketplace.InstallTool(*tool, packageManager)
	if err != nil {
		log.Fatal(err)
	}

}
