package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// OneLiner struct represents the structure of the one-liner YAML file
type OneLiner struct {
	ID   string `yaml:"id"`
	Info struct {
		Name        string `yaml:"name"`
		Author      string `yaml:"author"`
		Complexity  string `yaml:"complexity"`
		Description string `yaml:"description"`
	} `yaml:"info"`
	OneLiner struct {
		Cmd string `yaml:"cmd"`
	} `yaml:"one-liner"`
}

// LoadOneLiners loads all one-liners from YAML files in the specified directory
func LoadOneLiners(dir string) ([]OneLiner, error) {
	var oneLiners []OneLiner

	// filepath.Walk traverse récursivement le répertoire spécifié et tous ses sous-répertoires
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Vérifie que le fichier n'est pas un répertoire et a l'extension ".yaml"
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yml") {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			var oneLiner OneLiner
			// Unmarshal (analyse) le contenu YAML dans la structure OneLiner
			err = yaml.Unmarshal(data, &oneLiner)
			if err != nil {
				return err
			}

			// Ajoute le OneLiner à la liste
			oneLiners = append(oneLiners, oneLiner)
		}
		return nil
	})

	return oneLiners, err
}

func main() {
	// Charger tous les one-liners depuis le répertoire "one-liners"
	oneLiners, err := LoadOneLiners("one-liners")
	if err != nil {
		fmt.Printf("Error loading one-liners: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Affichage des one-liners disponibles
		fmt.Println("Available One-Liners:")
		fmt.Println("--------------------------------------------")
		for _, ol := range oneLiners {
			fmt.Printf("%s - %s\n", ol.ID, ol.Info.Name)
		}
		fmt.Println("--------------------------------------------")

		// Lecture de l'entrée utilisateur
		fmt.Print("H3cate4rsenal> ")
		input, err := reader.ReadString('\n') // Lire jusqu'à ce que l'utilisateur appuie sur Entrée
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		// Supprimer les espaces et nouvelles lignes autour de l'entrée
		input = strings.TrimSpace(input)

		// Si l'utilisateur entre "exit", sortir de la boucle
		if strings.ToLower(input) == "exit" {
			break
		}

		// Si l'utilisateur entre une commande "run <ID>", exécuter le one-liner
		if strings.HasPrefix(input, "run ") {
			oneLinerID := strings.TrimSpace(strings.TrimPrefix(input, "run "))
			found := false
			for _, ol := range oneLiners {
				if ol.ID == oneLinerID {
					runOneLiner(ol.OneLiner.Cmd) // Exécuter le one-liner correspondant
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("One-liner '%s' not found.\n", oneLinerID)
			}
		} else {
			fmt.Println("Invalid command. Please use 'run <ID>' or 'exit'.")
		}
	}
}

// runOneLiner executes the command associated with the one-liner
func runOneLiner(cmd string) {
	fmt.Printf("Running: %s\n", cmd)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}
	fmt.Printf("Output:\n%s\n", string(output))
}
