package one_liners

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// LoadOneLiners parcourt tous les sous-dossiers et charge les fichiers YAML comme one-liners
func LoadOneLiners(directory string) ([]OneLiner, error) {
	var oneLiners []OneLiner

	// Fonction interne récursive pour parcourir les répertoires
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Si l'élément n'est pas un fichier, continuer
		if info.IsDir() {
			return nil
		}

		// Lire le contenu du fichier
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", path, err)
			return nil
		}

		// Parse le contenu YAML du fichier dans la structure OneLiner
		var oneLiner OneLiner
		err = yaml.Unmarshal(content, &oneLiner)
		if err != nil {
			fmt.Printf("Failed to parse YAML in file %s: %v\n", path, err)
			return nil
		}

		// Ajouter le one-liner à la liste
		oneLiners = append(oneLiners, oneLiner)

		return nil
	})

	// Retourner une erreur si le parcours échoue
	if err != nil {
		return nil, err
	}

	return oneLiners, nil
}
