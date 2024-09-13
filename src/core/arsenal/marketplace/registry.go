package marketplace

import (
	"encoding/json"
	"io"
	"os"
)

// LoadRegistry charge la liste des outils depuis un fichier JSON.
func LoadRegistry(filepath string) (*Registry, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var registry Registry
	json.Unmarshal(byteValue, &registry)

	return &registry, nil
}

// FindTool recherche un outil par son nom dans le registry.
func (r *Registry) FindTool(name string) (*Tool, error) {
	for _, tool := range r.Tools {
		if tool.Name == name {
			return &tool, nil
		}
	}
	return nil, nil // ou renvoyer une erreur personnalisée si l'outil n'est pas trouvé
}
