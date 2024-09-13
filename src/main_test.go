package main

import (
	"testing"

	mp "github.com/Orangiuss/hecate/src/core/arsenal/marketplace"
	ol "github.com/Orangiuss/hecate/src/core/one-liners"
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
