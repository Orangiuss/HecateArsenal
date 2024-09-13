package one_liners

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func ShowOneLiners(oneLiners []OneLiner) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Nom", "Description", "Catégorie", "Commande"})

	if len(oneLiners) == 0 {
		fmt.Println("Aucun one-liner trouvé.")
		return
	}

	for _, ol := range oneLiners {
		// Vérifier le filtre si spécifié (par nom, catégorie ou tag)
		// Print test
		// fmt.Println(ol.Info.Name)
		// Ajouter les informations du one-liner dans la table
		table.Append([]string{
			ol.ID,
			ol.Info.Name,
			ol.Info.Description,
			ol.Info.Category,
			ol.OneLiner.Cmd,
		})
	}

	// Afficher la table
	table.Render()
}
