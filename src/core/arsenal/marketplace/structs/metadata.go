package marketplace_structs

// ToolMetadata représente les métadonnées associées à un outil.
type Metadata struct {
	Author      string `json:"author"`
	Version     string `json:"version"`
	Description string `json:"description"`
	License     string `json:"license"`
}
