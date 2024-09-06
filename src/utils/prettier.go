package utils

import "fmt"

// printAscii est une fonction qui permet d'afficher un texte en ASCII, avec une couleur spécifique, variable oui ou non, default oui, couleur par défaut rouge
func printAscii3(color string, variable bool) {
	text := `
	_   _                _          _                              _ 
	| | | | ___  ___ __ _| |_ ___   / \\   _ __ ___  ___ _ __   __ _| |
	| |_| |/ _ \/ __/ _` + "`" + ` | __/ _ \/ _ \ | '__/ __|/ _ \ '_ \ / _` + "`" + ` | |
	|  _  |  __/ (_| (_| | ||  __// ___ \| |  \__ \  __/ | | | (_| | |
	|_| |_|\___|\___\__,_|\__\___/_/   \_\_|  |___/\___|_| |_|\__,_|_|
																	  
	`

	if variable {
		fmt.Printf("\033[31m%s\033[0m", text)
	} else {
		fmt.Printf("\033[%sm%s\033[0m", color, text)
	}
}
