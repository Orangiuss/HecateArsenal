package main

import (
	"fmt"

	one_liners "github.com/Orangiuss/hecate/src/core/one-liners"
	utils_hecate "github.com/Orangiuss/hecate/src/core/utils"
)

func main() {
	utils_hecate.PrintAscii("32", false)

	oneLiners, err := one_liners.LoadOneLiners("one-liners-templates")
	if err != nil {
		fmt.Printf("Erreur de chargement des one-liners: %v\n", err)
		return
	}
	// Afficher le nombre de one-liners chargés
	fmt.Printf("Nombre de one-liners chargés: %d\n", len(oneLiners))

	// Afficher tous les one-liners disponibles
	one_liners.ShowOneLiners(oneLiners)

	// Exécuter un one-liner avec des variables
	vars := map[string]string{
		"DOMAIN":       "example.com",
		"STATUS_CODES": "200,301,302",
	}
	one_liners.RunOneLiner(oneLiners[0].OneLiner.Cmd, vars)

	data := "example"
	fmt.Println("MD5:", utils_hecate.HashMD5(data))
	fmt.Println("SHA256:", utils_hecate.HashSHA256(data))

	hexData := utils_hecate.ToHex(data)
	fmt.Println("To Hex:", hexData)
	decodedHex, _ := utils_hecate.FromHex(hexData)
	fmt.Println("From Hex:", decodedHex)

	hexDump := utils_hecate.ToHexDump(data)
	fmt.Println("To Hex Dump:", hexDump)
	decodedHexDump, _ := utils_hecate.FromHexDump(hexDump)
	fmt.Println("From Hex Dump:", decodedHexDump)

	match, _ := utils_hecate.RegexMatch(`[a-z]+`, data)
	fmt.Println("Regex Match:", match)

	replaced, _ := utils_hecate.RegexReplace(`e`, data, "E")
	fmt.Println("Regex Replace:", replaced)
}
