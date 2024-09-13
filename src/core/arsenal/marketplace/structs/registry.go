package marketplace_structs

// Registry gère la liste des outils disponibles dans le marketplace.
type Registry struct {
	Tools []Tool `json:"tools"`
}
