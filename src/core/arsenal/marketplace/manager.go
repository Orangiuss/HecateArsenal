package marketplace

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// // ManageTool exécute une ou plusieurs commandes pour une action donnée (install, remove, update) d'un outil en fonction du gestionnaire de paquets.
// func ManageTool(tool Tool, packageManager, actionType string) error {
// 	if pmActions, exists := tool.Actions[packageManager]; exists {
// 		var commands []string
// 		switch actionType {
// 		case "install":
// 			commands = pmActions.Install
// 		case "remove":
// 			commands = pmActions.Remove
// 		case "update":
// 			commands = pmActions.Update
// 		default:
// 			return fmt.Errorf("unknown action type: %s", actionType)
// 		}

// 		if len(commands) == 0 {
// 			return fmt.Errorf("no %s actions defined for tool: %s with package manager: %s", actionType, tool.Name, packageManager)
// 		}

// 		for _, cmd := range commands {
// 			if err := runCommand(cmd, tool.Name, actionType); err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	}
// 	return fmt.Errorf("no actions found for package manager: %s", packageManager)
// }

// InstallTool exécute toutes les commandes d'installation d'un outil en fonction du gestionnaire de paquets.
func InstallTool(tool Tool, packageManager string) error {
	if pmActions, exists := tool.Actions[packageManager]; exists {
		for _, cmd := range pmActions.Install {
			if err := runCommand(cmd, tool.Name, "install"); err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("no install actions defined for tool: %s with package manager: %s", tool.Name, packageManager)
}

// RemoveTool exécute toutes les commandes de suppression d'un outil en fonction du gestionnaire de paquets.
func RemoveTool(tool Tool, packageManager string) error {
	if pmActions, exists := tool.Actions[packageManager]; exists {
		for _, cmd := range pmActions.Remove {
			if err := runCommand(cmd, tool.Name, "remove"); err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("no remove actions defined for tool: %s with package manager: %s", tool.Name, packageManager)
}

// UpdateTool exécute toutes les commandes de mise à jour d'un outil en fonction du gestionnaire de paquets.
func UpdateTool(tool Tool, packageManager string) error {
	if pmActions, exists := tool.Actions[packageManager]; exists {
		for _, cmd := range pmActions.Update {
			if err := runCommand(cmd, tool.Name, "update"); err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("no update actions defined for tool: %s with package manager: %s", tool.Name, packageManager)
}

// TestTool teste si l'outil est installé en exécutant les commandes de test définies.
func TestTool(tool Tool) error {
	for _, command := range tool.TestCommands {
		cmd := exec.Command("sh", "-c", command) // Exécute la commande dans un shell
		output, err := cmd.CombinedOutput()      // Combine stdout et stderr
		if err != nil {
			return fmt.Errorf("error running test command for tool %s: %s\nOutput: %s", tool.Name, err, output)
		}

		// Vérifie si la commande a produit une sortie et l'affiche pour le débogage
		fmt.Printf("Test command for tool %s:\n%s\n", tool.Name, output)
	}
	return nil
}

// runCommand exécute une commande shell spécifique pour un outil et une action donnée.
func runCommand(cmdStr, toolName, action string) error {
	fmt.Printf("Running %s command for tool %s: %s\n", action, toolName, cmdStr)
	// Exécution de la commande shell
	cmd := exec.Command("sh", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running %s command for tool %s: %v\nOutput: %s", action, toolName, err, string(output))
	}
	fmt.Printf("Command output: %s\n", string(output))
	return nil
}

// ShowTools affiche les outils disponibles dans le registry.
func ShowTools(r *Registry) {
	tools, _ := r.GetTools()
	ShowToolsTools(tools)
}

// ShowToolsByCategory affiche les outils par catégorie dans le registry.
func ShowToolsByCategory(r *Registry, category string) {
	tools := r.FindToolsByCategory(category)
	if len(tools) == 0 {
		fmt.Printf("Aucun outil trouvé pour la catégorie: %s\n", category)
		return
	}
	ShowTools(&Registry{Tools: tools})
}

// ShowToolsByTag affiche les outils par tag dans le registry.
func ShowToolsByTag(r *Registry, tag string) {
	tools := r.FindToolsByTag(tag)
	if len(tools) == 0 {
		fmt.Printf("Aucun outil trouvé pour le tag: %s\n", tag)
		return
	}
	ShowTools(&Registry{Tools: tools})
}

// ShowToolsByAuthor affiche les outils par auteur dans le registry.
func ShowToolsByAuthor(r *Registry, author string) {
	tools := r.FindToolsByAuthor(author)
	if len(tools) == 0 {
		fmt.Printf("Aucun outil trouvé pour l'auteur: %s\n", author)
		return
	}
	ShowTools(&Registry{Tools: tools})
}

// ShowToolsTools affiche les outils, sans passer par le registry.
func ShowToolsTools(tools []Tool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Description", "Version", "Author"})

	if len(tools) == 0 {
		fmt.Println("Aucun outil trouvé.")
		return
	}

	for _, tool := range tools {
		commands := ""
		for _, cmd := range tool.TestCommands {
			commands += cmd + "; "
		}

		// Ajouter les informations de l'outil dans la table
		table.Append([]string{
			tool.ID,
			tool.Name,
			tool.Description,
			tool.Version,
			tool.Author,
		})
	}

	// Afficher la table
	table.Render()
}

// FindTool recherche des outils dont le nom contient la sous-chaîne spécifiée dans le registre.
func (r *Registry) FindTool(query string) ([]Tool, error) {
	var results []Tool

	for _, tool := range r.Tools {
		if strings.Contains(strings.ToLower(tool.Name), strings.ToLower(query)) {
			results = append(results, tool)
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("aucun outil trouvé pour la requête: %s", query)
	}

	return results, nil
}

// FindToolExact recherche un outil par nom exact dans le registry.
func (r *Registry) FindToolExact(name string) (*Tool, error) {
	for _, tool := range r.Tools {
		if tool.Name == name {
			return &tool, nil
		}
	}
	return nil, fmt.Errorf("aucun outil trouvé pour le nom: %s", name)
}

// FindToolByID recherche un outil par ID dans le registry.
func (r *Registry) FindToolByID(id string) (*Tool, error) {
	for _, tool := range r.Tools {
		if tool.ID == id {
			return &tool, nil
		}
	}
	return nil, fmt.Errorf("aucun outil trouvé pour l'ID: %s", id)
}

// FindToolsByAuthor recherche les outils par auteur dans le registry.
func (r *Registry) FindToolsByAuthor(author string) []Tool {
	var results []Tool
	for _, tool := range r.Tools {
		if tool.Metadata.Author == author {
			results = append(results, tool)
		}
	}
	return results
}

// FindToolsByCategory recherche les outils par catégorie dans le registry.
func (r *Registry) FindToolsByCategory(category string) []Tool {
	var results []Tool
	for _, tool := range r.Tools {
		for _, cat := range tool.Categories {
			if cat == category {
				results = append(results, tool)
				break
			}
		}
	}
	return results
}

// FindToolsByTag recherche les outils par tag
func (r *Registry) FindToolsByTag(tag string) []Tool {
	var results []Tool
	for _, tool := range r.Tools {
		for _, t := range tool.Tags {
			if t == tag {
				results = append(results, tool)
				break
			}
		}
	}
	return results
}

// ShowToolById affiche les informations détaillées sur un outil spécifique.
func ShowToolById(r *Registry, toolID string) {
	tool, err := r.FindToolByID(toolID)
	if err != nil {
		fmt.Println(err)
		return
	}
	ShowTool(*tool)
}

// ShowTool affiche les informations détaillées sur un outil spécifique.
func ShowTool(tool Tool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})

	// Ajouter les détails de l'outil dans la table
	table.Append([]string{"ID", tool.ID})
	table.Append([]string{"Name", tool.Name})
	table.Append([]string{"Description", tool.Description})
	table.Append([]string{"Version", tool.Version})
	table.Append([]string{"Author", tool.Author})
	table.Append([]string{"Commands", formatCommands(tool.TestCommands)})
	table.Append([]string{"Categories", formatList(tool.Categories)})
	table.Append([]string{"Tags", formatList(tool.Tags)})
	table.Append([]string{"Metadata", formatMetadata(tool.Metadata)})

	// Afficher la table
	table.Render()
}

// formatCommands formate les commandes pour l'affichage.
func formatCommands(commands []string) string {
	if len(commands) == 0 {
		return "N/A"
	}
	return formatList(commands)
}

// formatList formate une liste pour l'affichage.
func formatList(list []string) string {
	if len(list) == 0 {
		return "N/A"
	}
	return fmt.Sprintf("[%s]", fmt.Sprintf("%s", list))
}

// formatMetadata formate les métadonnées pour l'affichage.
func formatMetadata(metadata Metadata) string {
	if (metadata == Metadata{}) {
		return "N/A"
	}
	return fmt.Sprintf("License: %s; Source: %s; Logo URL: %s", metadata.License, metadata.Source, metadata.Logo_url)
}
