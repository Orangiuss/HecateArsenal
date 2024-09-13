package marketplace

import "fmt"

// GetActionByPackageManager récupère les actions associées à un gestionnaire de paquets spécifique.
func (t *Tool) GetActionByPackageManager(manager string, action string) ([]string, error) {
	if pmActions, exists := t.Actions[manager]; exists {
		switch action {
		case "install":
			return pmActions.Install, nil
		case "remove":
			return pmActions.Remove, nil
		case "update":
			return pmActions.Update, nil
		}
	}
	return nil, fmt.Errorf("action %s not found for package manager %s", action, manager)
}
