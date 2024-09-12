package main

import (
	"fmt"

	one_liners "github.com/Orangiuss/hecate/src/core/one-liners"

	"github.com/Orangiuss/hecate/src/core/plan_organiser"
)

func main() {
	// Créer une nouvelle instance de PlanOrganiser
	plan := plan_organiser.NewPlanOrganiser()

	// Créer un OneLiner d'exemple
	exampleOneLiner := one_liners.OneLiner{
		ID: "example-1",
		Info: one_liners.Info{
			Name:        "Nmap Scan",
			Description: "Scan réseau basique avec Nmap",
			Category:    "Reconnaissance",
			Tags:        []string{"nmap", "reconnaissance"},
		},
		OneLiner: one_liners.OneLinerCmd{
			Cmd: "nmap -sP 192.168.0.0/24",
		},
	}

	// Ajouter le OneLiner à la phase de reconnaissance
	plan.AddOneLiner(exampleOneLiner, "reconnaissance")

	// Sauvegarder le plan
	err := plan.SavePlan("mon-plan.yaml")
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde du plan:", err)
		return
	}

	// Charger le plan à partir d'un fichier
	loadedPlan, err := plan_organiser.LoadPlan("mon-plan.yaml")
	if err != nil {
		fmt.Println("Erreur lors du chargement du plan:", err)
		return
	}

	// Exécuter le plan
	loadedPlan.RunPlan()

	// Afficher le plan chargé
	fmt.Printf("Plan chargé : %+v\n", loadedPlan)
}
