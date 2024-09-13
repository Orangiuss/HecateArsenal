package marketplace

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func NewRegistry() *Registry {
	return &Registry{
		Tools: []Tool{}, // Assurez-vous que Tools est initialisé, même s'il est vide
	}
}

// LoadRegistry charge la liste des outils depuis un fichier JSON.
func LoadRegistry(filepath string) (*Registry, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var registry = NewRegistry()
	json.Unmarshal(byteValue, &registry)

	return registry, nil
}

// Get all tools from the registry
func (r *Registry) GetTools() ([]Tool, error) {
	if r == nil {
		fmt.Println("Registry is not initialized properly or is empty.")
		os.Exit(1)
	}
	return r.Tools, nil
}
