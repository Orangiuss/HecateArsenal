package main

import (
	"fmt"
	"testing"

	mp "github.com/Orangiuss/hecate/src/core/arsenal/marketplace"
	ol "github.com/Orangiuss/hecate/src/core/one-liners"
	"github.com/Orangiuss/hecate/src/core/plan_organiser"
)

// TestRunOneLiner teste la fonction RunOneLiner avec une commande fictive
func TestRunOneLiner(t *testing.T) {
	exampleOneLiner := ol.OneLiner{
		ID: "example-1",
		Info: ol.Info{
			Name:        "Nmap Scan",
			Description: "Scan réseau basique avec Nmap",
			Category:    "Reconnaissance",
			Tags:        []string{"nmap", "reconnaissance"},
		},
		OneLiner: ol.OneLinerCmd{
			Cmd: "nmap -V",
		},
	}
	variables := map[string]string{}

	// Simulez la commande ou utilisez un package de simulation
	err := ol.RunOneLiner(exampleOneLiner, variables)
	if err != nil {
		t.Errorf("RunOneLiner() failed: %v", err)
	}
}

func TestInstallTool(t *testing.T) {
	tool := mp.Tool{
		Name: "nmap",
		Actions: map[string]mp.PackageManagerActions{
			"apt": {
				Install: []string{"sudo apt-get install -y nmap"},
				Remove:  []string{"sudo apt-get remove -y nmap"},
			},
		},
	}

	// Simulez le package manager ou le résultat de la commande
	err := mp.InstallTool(tool, "apt")
	if err != nil {
		t.Errorf("InstallTool() failed: %v", err)
	}
}

func TestPlanOrganiser(t *testing.T) {
	// Créer une nouvelle instance de PlanOrganiser
	plan := plan_organiser.NewPlanOrganiser()

	// Créer un OneLiner d'exemple
	exampleOneLiner := ol.OneLiner{
		ID: "example-1",
		Info: ol.Info{
			Name:        "Nmap Scan",
			Description: "Scan réseau basique avec Nmap",
			Category:    "Reconnaissance",
			Tags:        []string{"nmap", "reconnaissance"},
		},
		OneLiner: ol.OneLinerCmd{
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

func TestGetActionByPackageManager(t *testing.T) {
	tool := mp.Tool{
		Actions: map[string]mp.PackageManagerActions{
			"apt": {
				Install: []string{"sudo apt-get install"},
				Remove:  []string{"sudo apt-get remove"},
			},
		},
	}

	installCmds, err := tool.GetActionByPackageManager("apt", "install")
	if err != nil {
		t.Errorf("GetActionByPackageManager() failed: %v", err)
	}

	// Vérifiez les commandes install
	expectedCmd := "sudo apt-get install"
	if len(installCmds) == 0 || installCmds[0] != expectedCmd {
		t.Errorf("GetActionByPackageManager() = %v; want %v", installCmds, expectedCmd)
	}
}
