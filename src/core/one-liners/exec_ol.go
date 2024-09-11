package one_liners

import (
	"fmt"
	"os/exec"
	"strings"
)

// RunOneLiner exécute une commande one-liner avec les variables fournies
func RunOneLiner(cmdTemplate string, variables map[string]string) error {
	// Remplacer les variables dans le template de commande
	for key, value := range variables {
		placeholder := fmt.Sprintf("{%s}", strings.ToUpper(key))
		cmdTemplate = strings.ReplaceAll(cmdTemplate, placeholder, value)
	}

	// Exécuter la commande
	cmd := exec.Command("bash", "-c", cmdTemplate)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution de la commande : %v\n%s", err, output)
	}

	// Afficher le résultat de la commande
	fmt.Println(string(output))
	return nil
}
