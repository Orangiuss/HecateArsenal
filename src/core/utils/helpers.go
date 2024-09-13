package utils

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/mbndr/figlet4go"
)

func IsRoot() bool {
	usr, err := user.Current()
	if err != nil {
		return false
	}
	return usr.Uid == "0"
}

func PrintHelloWorld() {
	printColor("Hello, World!", 32)
}

func printColor(s string, n int) {
	fmt.Println("\033[" + strconv.Itoa(n) + "m" + s + "\033[0m")
}

func PrintAscii(color string, variable bool) {
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
