package marketplace_structs

// PackageManagerActions représente les actions pour un gestionnaire de paquets spécifique
type PackageManagerActions struct {
	Install string `json:"install"`
	Remove  string `json:"remove"`
	Update  string `json:"update"`
}
