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
