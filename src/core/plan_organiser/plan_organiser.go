package plan_organiser

import (
	"os"

	ol "github.com/Orangiuss/hecate/src/core/one-liners"
	"gopkg.in/yaml.v2"
)

// AddOneLiner ajoute un one-liner à une étape spécifique du plan d'attaque
func (p *PlanOrganiser) AddOneLiner(oneLiner ol.OneLiner, step string) {
	switch step {
	case "reconnaissance":
		p.Reconnaissance = append(p.Reconnaissance, oneLiner)
	case "enumeration":
		p.Enumeration = append(p.Enumeration, oneLiner)
	case "exploitation":
		p.Exploitation = append(p.Exploitation, oneLiner)
	case "post-exploitation":
		p.PostExploitation = append(p.PostExploitation, oneLiner)
	case "rapport":
		p.Rapport = append(p.Rapport, oneLiner)
	}
}

// RemoveOneLiner supprime un one-liner d'une étape spécifique du plan d'attaque
func (p *PlanOrganiser) RemoveOneLiner(id string, step string) {
	switch step {
	case "reconnaissance":
		p.Reconnaissance = removeOneLiner(p.Reconnaissance, id)
	case "enumeration":
		p.Enumeration = removeOneLiner(p.Enumeration, id)
	case "exploitation":
		p.Exploitation = removeOneLiner(p.Exploitation, id)
	case "post-exploitation":
		p.PostExploitation = removeOneLiner(p.PostExploitation, id)
	case "rapport":
		p.Rapport = removeOneLiner(p.Rapport, id)
	}
}

// removeOneLiner supprime un one-liner d'une liste de one-liners en fonction de son ID
func removeOneLiner(oneLiners []ol.OneLiner, id string) []ol.OneLiner {
	var updatedOneLiners []ol.OneLiner
	for _, oneLiner := range oneLiners {
		if oneLiner.ID != id {
			updatedOneLiners = append(updatedOneLiners, oneLiner)
		}
	}
	return updatedOneLiners
}

// SavePlan enregistre le plan d'attaque dans un fichier YAML
func (p *PlanOrganiser) SavePlan(filename string) error {
	data, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadPlan charge un plan d'attaque à partir d'un fichier YAML
func LoadPlan(filename string) (*PlanOrganiser, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var plan PlanOrganiser
	err = yaml.Unmarshal(data, &plan)
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

// RunPlan exécute le plan d'attaque en fonction des étapes
func (p *PlanOrganiser) RunPlan() {
	// Reconnaissance
	for _, oneLiner := range p.Reconnaissance {
		variablesMap := ol.ConvertVariablesToMap(oneLiner.Variables)
		ol.RunOneLiner(oneLiner, variablesMap)
	}

	// Enumeration
	for _, oneLiner := range p.Enumeration {
		variablesMap := ol.ConvertVariablesToMap(oneLiner.Variables)
		ol.RunOneLiner(oneLiner, variablesMap)
	}

	// Exploitation
	for _, oneLiner := range p.Exploitation {
		variablesMap := ol.ConvertVariablesToMap(oneLiner.Variables)
		ol.RunOneLiner(oneLiner, variablesMap)
	}

	// Post-exploitation
	for _, oneLiner := range p.PostExploitation {
		variablesMap := ol.ConvertVariablesToMap(oneLiner.Variables)
		ol.RunOneLiner(oneLiner, variablesMap)
	}

	// Rapport
	for _, oneLiner := range p.Rapport {
		variablesMap := ol.ConvertVariablesToMap(oneLiner.Variables)
		ol.RunOneLiner(oneLiner, variablesMap)
	}
}
