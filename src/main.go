package main

import (
	"fmt"
	one_liners "hecatearsenal/one-liners"
	"hecatearsenal/plan_organiser"
	"hecatearsenal/utils"
)

func main() {
	// Appeler une fonction du package one_liners
	oneLiners, err := one_liners.LoadOneLiners("one-liners")
	if err != nil {
		fmt.Printf("Error loading one-liners: %v\n", err)
		return
	}

	// Appeler une fonction du package utils
	utils.PrintHelloWorld()

	// Affichage des one-liners
	for _, ol := range oneLiners {
		fmt.Printf("One-Liner: %s\n", ol.Info.Name)
	}

	// Afficher les one-liners dans une table
	one_liners.ShowOneLiners(oneLiners, "", "")
	PlanOrganiser := plan_organiser.NewPlanOrganiser()
}
