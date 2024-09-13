package marketplace

type Tool struct {
	ID           string                           `json:"id"`
	Name         string                           `json:"name"`
	Description  string                           `json:"description"`
	Version      string                           `json:"version"`
	Author       string                           `json:"author"`
	TestCommands []string                         `json:"test-commands"`
	Actions      map[string]PackageManagerActions `json:"actions"`    // Actions spécifiques à chaque gestionnaire de paquets
	Categories   []string                         `json:"categories"` // Nouveaux champs pour catégoriser les outils
	Tags         []string                         `json:"tags"`       // Tags pour filtrer les outils
	Metadata     Metadata                         `json:"metadata"`
}

// ToolMetadata représente les métadonnées associées à un outil.
type Metadata struct {
	Author      string `json:"author"`
	Version     string `json:"version"`
	Description string `json:"description"`
	License     string `json:"license"`
	Logo_url    string `json:"logo-url"`
}

// PackageManagerActions représente les actions pour un gestionnaire de paquets spécifique
type PackageManagerActions struct {
	Install []string `json:"install"`
	Remove  []string `json:"remove"`
	Update  []string `json:"update"`
}

// Registry gère la liste des outils disponibles dans le marketplace.
type Registry struct {
	Tools []Tool `json:"tools"`
}
