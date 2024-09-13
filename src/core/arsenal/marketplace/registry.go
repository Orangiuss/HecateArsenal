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

// Get all tools from the registry
func (r *Registry) GetTools() ([]Tool, error) {
	return r.Tools, nil
}
