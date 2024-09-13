package appdata

import (
	"fmt"
	"log"
	"os"

	mp "github.com/Orangiuss/hecate/src/core/arsenal/marketplace"
	ol "github.com/Orangiuss/hecate/src/core/one-liners"
)

var Version = "0.0.1"
var Author = "Orangiuss"
var oneLiners []ol.OneLiner
var tools []mp.Tool
var registry *mp.Registry

// GetDataPath récupère le chemin du répertoire de données, soit par défaut, soit par une variable d'environnement.
func GetDataPath() string {
	if path := os.Getenv("HECATE_DATA_PATH"); path != "" {
		return path
	}
	return "./data" // Chemin par défaut
}

// Initialiser les données
func Init() error {

	files := []string{
		"one-liners-templates",
		"tools.json",
	}
	dataPath := GetDataPath()
	for _, file := range files {
		fullPath := fmt.Sprintf("%s/%s", dataPath, file)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			return fmt.Errorf("le fichier/dossier %s est manquant à %s: %v", file, fullPath, err)
		}
	}

	var err error
	oneLiners, err = ol.LoadOneLiners("one-liners-templates")
	if err != nil {
		return err
	}

	// Charger le registre des outils
	registry, err := mp.LoadRegistry("src/core/arsenal/tools.json")
	if err != nil {
		log.Fatal(err)
	}

	// Charger les outils depuis le registre
	tools, err = registry.GetTools()

	// Check if there is an error
	if err != nil {
		return fmt.Errorf("erreur lors du chargement du registre des outils : %v", err)
	}

	return nil
}

// Récupérer les one-liners
func GetOneLiners() []ol.OneLiner {
	return oneLiners
}

// Récupérer les tools
func GetTools() []mp.Tool {
	return tools
}

// Récupérer le registry
func GetRegistry() *mp.Registry {
	return registry
}
