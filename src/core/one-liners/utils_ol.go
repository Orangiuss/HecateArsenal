package one_liners

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
