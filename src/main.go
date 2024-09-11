package main

import (
	"fmt"
	one_liners "hecate/core/one-liners"
	utils_hecate "hecate/core/utils"
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

	data := "example&%$"
	fmt.Println("MD5:", utils_hecate.HashMD5(data))
	fmt.Println("SHA256:", utils_hecate.HashSHA256(data))
	fmt.Println("NTLM:", utils_hecate.HashNTLM(data))

	bcryptHash, _ := utils_hecate.HashBcrypt(data)
	fmt.Println("Bcrypt:", bcryptHash)

	base64Encoded := utils_hecate.EncodeBase64(data)
	fmt.Println("Base64 Encoded:", base64Encoded)
	base64Decoded, _ := utils_hecate.DecodeBase64(base64Encoded)
	fmt.Println("Base64 Decoded:", base64Decoded)

	urlEncoded := utils_hecate.URLEncode(data)
	fmt.Println("URL Encoded:", urlEncoded)
	urlDecoded, _ := utils_hecate.URLDecode(urlEncoded)
	fmt.Println("URL Decoded:", urlDecoded)
}
