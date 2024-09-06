package plan_organiser

// PlanOrganiser est une structure qui contient une liste de méthodologies pour le pentest, avec les one-liners du module one_liners
// Nous allons avoir les étapes suivantes:
// 1. Reconnaissance
// 2. Enumération
// 3. Exploitation
// 4. Post-exploitation
// 5. Rapport

// l'utilisateur peut définir son plan d'attaque en ajoutant ou en supprimant des one-liners à chaque étape, nous allons avoir des talbeaux
// de one-liners pour chaque étape, et l'utilisateur peut les manipuler comme il le souhaite, et les sauver dans un fichier YAML
// pour les réutiliser plus tard, avec l'ID de chaque one-liner comme clé.
type PlanOrganiser struct {
	Reconnaissance   []OneLiner
	Enumeration      []OneLiner
	Exploitation     []OneLiner
	PostExploitation []OneLiner
	Rapport          []OneLiner
}

// NewPlanOrganiser crée une nouvelle instance de PlanOrganiser avec des tableaux vides pour chaque étape
func NewPlanOrganiser() *PlanOrganiser {
	return &PlanOrganiser{
		Reconnaissance:   []OneLiner{},
		Enumeration:      []OneLiner{},
		Exploitation:     []OneLiner{},
		PostExploitation: []OneLiner{},
		Rapport:          []OneLiner{},
	}
}
