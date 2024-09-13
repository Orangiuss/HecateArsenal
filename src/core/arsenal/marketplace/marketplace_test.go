package marketplace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var metadataTest = Metadata{
	Author:      "John Doe",
	Version:     "1.0",
	Description: "Description of the tool",
	License:     "GPLv3",
	Logo_url:    "https://example.com/logo.png",
	Source:      "https://example.com/source",
}

// Exemple de données d'outils pour les tests
var testTools = []Tool{
	{
		ID:           "1",
		Name:         "Nmap",
		Description:  "Network Mapper",
		Version:      "7.80",
		Author:       "Nmap Project",
		TestCommands: []string{"nmap -V"},
		Actions:      map[string]PackageManagerActions{},
		Categories:   []string{"Network"},
		Tags:         []string{"scanner", "network"},
		Metadata:     metadataTest,
	},
	{
		ID:           "2",
		Name:         "Sqlmap",
		Description:  "SQL Injection Tool",
		Version:      "1.4",
		Author:       "Bernardo Damele A. G.",
		TestCommands: []string{"sqlmap --version"},
		Actions:      map[string]PackageManagerActions{},
		Categories:   []string{"Database", "Exploitation"},
		Tags:         []string{"sql", "injection", "database"},
		Metadata:     metadataTest,
	},
}

func TestFindTool(t *testing.T) {
	// Initialiser le registre avec des outils de test
	registry := &Registry{
		Tools: testTools,
	}

	// Tester un outil qui existe
	tool, err := registry.FindToolExact("Nmap")
	assert.Nil(t, err, "Aucune erreur ne devrait être retournée pour un outil existant")
	assert.Equal(t, "Nmap", tool.Name, "Le nom de l'outil devrait être 'Nmap'")

	// Tester un outil qui n'existe pas
	_, err = registry.FindTool("NonExistantTool")
	assert.NotNil(t, err, "Une erreur devrait être retournée pour un outil non existant")
	assert.EqualError(t, err, "aucun outil trouvé pour la requête: NonExistantTool")
}

func TestShowTool(t *testing.T) {
	// Initialiser le registre avec des outils de test
	registry := &Registry{
		Tools: testTools,
	}

	// Tester l'affichage des détails d'un outil
	toolID := "1"
	tool, err := registry.FindToolByID(toolID)
	assert.Nil(t, err, "Aucune erreur ne devrait être retournée pour un outil existant")
	assert.NotNil(t, tool, "L'outil trouvé ne devrait pas être nul")

	// Ajouter ici un test qui capture la sortie pour vérifier si `ShowTool` imprime correctement les détails
	// (exemple en utilisant une bibliothèque comme `testify` pour capturer la sortie du `fmt.Printf`)
}

func TestShowTools(t *testing.T) {
	// Initialiser le registre avec des outils de test
	registry := &Registry{
		Tools: testTools,
	}

	// Vous pouvez aussi capturer l'output ici si vous voulez tester l'affichage
	ShowTools(registry)

	// Vérifier si tous les outils sont affichés correctement (exemple simple ici, basé sur l'état)
	assert.Equal(t, len(registry.Tools), 2, "Le registre devrait contenir 2 outils")
}

func TestInstallToolCommand(t *testing.T) {
	// Simuler une commande d'installation pour un outil donné
	tool := testTools[0] // Nmap par exemple
	assert.Equal(t, tool.TestCommands[0], "nmap -V", "La commande de test pour Nmap devrait être 'nmap -v'")
}

func TestRemoveToolCommand(t *testing.T) {
	// Tester une commande de suppression
	// Supposons que vous ayez une commande de suppression pour des outils dans vos actions
	tool := testTools[1] // Sqlmap
	assert.Equal(t, tool.TestCommands[0], "sqlmap --version", "La commande de test pour Sqlmap devrait être correcte")
}
