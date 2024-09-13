package one_liners

import (
	"regexp"
	"strings"
)

// ConvertVariablesToMap convertit un map de Variable en map de string
func ConvertVariablesToMap(vars map[string]Variable) map[string]string {
	result := make(map[string]string)
	for key, variable := range vars {
		if variable.Required {
			result[key] = variable.Default
		}
	}
	return result
}

// GetOneLinerByName retourne un one-liner en fonction de son nom
func GetOneLinerByName(oneLiners []OneLiner, name string) OneLiner {
	for _, ol := range oneLiners {
		if ol.Info.Name == name {
			return ol
		}
	}
	return OneLiner{}
}

// GetOneLinerByID retourne un one-liner en fonction de son ID
func GetOneLinerByID(oneLiners []OneLiner, id string) OneLiner {
	for _, ol := range oneLiners {
		if ol.ID == id {
			return ol
		}
	}
	return OneLiner{}
}

// SearchOneLiners recherche des one-liners en fonction d'un filtre de type LIKE utilisant des expressions régulières.
func SearchOneLiners(oneLiners []OneLiner, filter string) ([]OneLiner, error) {
	var result []OneLiner

	// Convertir le filtre de type LIKE en une expression régulière
	regexPattern := "^" + strings.ReplaceAll(filter, "%", ".*") + "$"
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, err
	}

	for _, ol := range oneLiners {
		if re.MatchString(ol.Info.Name) {
			result = append(result, ol)
		}
	}

	return result, nil
}

// GetOneLinersByCategory retourne une liste de one-liners en fonction de la catégorie
func GetOneLinersByCategory(oneLiners []OneLiner, category string) []OneLiner {
	var result []OneLiner
	for _, ol := range oneLiners {
		if ol.Info.Category == category {
			result = append(result, ol)
		}
	}
	return result
}

// GetOneLinersByTag retourne une liste de one-liners en fonction du tag
func GetOneLinersByTag(oneLiners []OneLiner, tag string) []OneLiner {
	var result []OneLiner
	for _, ol := range oneLiners {
		for _, t := range ol.Info.Tags {
			if t == tag {
				result = append(result, ol)
				break
			}
		}
	}
	return result
}
