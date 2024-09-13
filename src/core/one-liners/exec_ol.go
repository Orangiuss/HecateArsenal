package one_liners

import (
	"fmt"
	"os/exec"
	"strings"
)

// // RunOneLiner exécute une commande one-liner avec les variables fournies
// func RunOneLiner(cmdTemplate string, variables map[string]string) error {
// 	// Remplacer les variables dans le template de commande
// 	for key, value := range variables {
// 		placeholder := fmt.Sprintf("{%s}", strings.ToUpper(key))
// 		cmdTemplate = strings.ReplaceAll(cmdTemplate, placeholder, value)
// 	}

// 	// Exécuter la commande
// 	cmd := exec.Command("bash", "-c", cmdTemplate)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("erreur lors de l'exécution de la commande : %v\n%s", err, output)
// 	}

// 	// Afficher le résultat de la commande
// 	fmt.Println(string(output))
// 	return nil
// }

// RunOneLiner exécute une commande one-liner avec les variables fournies
func RunOneLiner(oneLiner OneLiner, variables map[string]string) error {
	// Remplacer les variables dans le template de commande
	cmdTemplate := oneLiner.OneLiner.Cmd

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

// TestOneLiner exécute une commande one-liner avec les variables fournies
// et retourne le résultat de la commande
func TestOneLiner() (string, error) {
	// Exemple d'initialisation
	oneLiner := OneLiner{
		ID: "example",
		Info: Info{
			Name:        "Example OneLiner",
			Description: "This is an example one-liner.",
			Category:    "Example",
			Author:      "Author",
			Complexity:  "Simple",
			Tags:        []string{"example", "one-liner"},
		},
		Variables: map[string]Variable{
			"HOST": {Description: "The host to connect to", Required: true},
			"PORT": {Description: "The port to connect to", Required: false, Default: "80"},
		},
		OneLiner: OneLinerCmd{
			Cmd: "echo Connecting to {HOST} on port {PORT}",
		},
	}

	// Exemple de variables
	variables := map[string]string{
		"HOST": "localhost",
		"PORT": "8080",
	}

	// Exécuter le one-liner
	err := RunOneLiner(oneLiner, variables)
	if err != nil {
		fmt.Println("Erreur:", err)
	}
	return "", nil
}
