package one_liners

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// LoadOneLiners charge tous les one-liners à partir du répertoire spécifié
func LoadOneLiners(directory string) ([]OneLiner, error) {
	// Lire les fichiers dans le répertoire en utilisant os.ReadDir
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var oneLiners []OneLiner

	// Parcourir chaque fichier dans le répertoire
	for _, file := range files {
		// Ignorer les sous-répertoires
		if file.IsDir() {
			continue
		}

		// Lire le contenu du fichier avec os.ReadFile
		content, err := os.ReadFile(directory + "/" + file.Name())
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", file.Name(), err)
			continue
		}

		// Parse le contenu YAML du fichier dans la structure OneLiner
		var oneLiner OneLiner
		err = yaml.Unmarshal(content, &oneLiner)
		if err != nil {
			fmt.Printf("Failed to parse YAML in file %s: %v\n", file.Name(), err)
			continue
		}

		// Utiliser le nom du fichier comme identifiant (ID) du one-liner
		oneLiner.ID = file.Name()

		// Ajouter le one-liner au tableau
		oneLiners = append(oneLiners, oneLiner)
	}

	return oneLiners, nil
}
