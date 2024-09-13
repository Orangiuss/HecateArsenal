package marketplace_structs

type Tool struct {
	ID            string                           `json:"id"`
	Name          string                           `json:"name"`
	Description   string                           `json:"description"`
	Version       string                           `json:"version"`
	Author        string                           `json:"author"`
	TestsCommands []string                         `json:"test-commands"`
	Actions       map[string]PackageManagerActions `json:"actions"`    // Actions spécifiques à chaque gestionnaire de paquets
	Categories    []string                         `json:"categories"` // Nouveaux champs pour catégoriser les outils
	Tags          []string                         `json:"tags"`       // Tags pour filtrer les outils
	Metadata      Metadata                         `json:"metadata"`
}
