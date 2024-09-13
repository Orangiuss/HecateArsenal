package utils

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

// printAscii est une fonction qui permet d'afficher un texte en ASCII, avec une couleur spécifique, variable oui ou non, default oui, couleur par défaut rouge
func PrintAscii2(_ string) {
	text := "Hecate"
	ascii := figlet4go.NewAsciiRender()

	// Ajouter des options
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorRed,
	}

	// Rendu du texte
	renderStr, _ := ascii.RenderOpts(text, options)
	fmt.Println(renderStr)
}
