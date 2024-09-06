package utils

import (
	"fmt"
	"strconv"
)

func PrintHelloWorld() {
	printColor("Hello, World!", 32)
}

func printColor(s string, n int) {
	fmt.Println("\033[" + strconv.Itoa(n) + "m" + s + "\033[0m")
}
