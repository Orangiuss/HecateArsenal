package main

import (
	"fmt"
	one_liners "hecatearsenal/one-liners"
	utils_ha "hecatearsenal/utils"
)

func main() {
	// Appeler un fonction du package utils
	utils_ha.PrintHelloWorld()
	utils_ha.PrintAscii("32", false)

	oneLiners, err := one_liners.LoadOneLiners("one-liners-templates")
	if err != nil {
		fmt.Printf("Erreur de chargement des one-liners: %v\n", err)
		return
	}
	// Afficher le nombre de one-liners chargés
	fmt.Printf("Nombre de one-liners chargés: %d\n", len(oneLiners))

	// Afficher tous les one-liners disponibles
	one_liners.ShowOneLiners(oneLiners)

	// Exécuter un one-liner avec des variables
	vars := map[string]string{
		"DOMAIN":       "example.com",
		"STATUS_CODES": "200,301,302",
	}
	one_liners.RunOneLiner(oneLiners[0].OneLiner.Cmd, vars)
}
