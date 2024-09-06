package one_liners

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func ShowOneLiners(oneLiners []OneLiner, filter string, filterType string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Nom", "Description", "Catégorie", "Tags", "Commande"})

	for _, ol := range oneLiners {
		// Vérifier le filtre si spécifié (par nom, catégorie ou tag)
		switch filterType {
		case "name":
			if !strings.Contains(strings.ToLower(ol.Info.Name), strings.ToLower(filter)) {
				continue
			}
		case "category":
			if !strings.EqualFold(strings.ToLower(ol.Info.Category), strings.ToLower(filter)) {
				continue
			}
		case "tag":
			tagMatch := false
			for _, tag := range ol.Info.Tags {
				if strings.EqualFold(strings.ToLower(tag), strings.ToLower(filter)) {
					tagMatch = true
					break
				}
			}
			if !tagMatch {
				continue
			}
		}

		// Ajouter les informations du one-liner dans la table
		table.Append([]string{
			ol.ID,
			ol.Info.Name,
			ol.Info.Description,
			ol.Info.Category,
			strings.Join(ol.Info.Tags, ", "),
			ol.OneLiner.Cmd,
		})
	}

	// Afficher la table
	table.Render()
}
