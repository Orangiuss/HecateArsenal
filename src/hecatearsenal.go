package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
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
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
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
	// Load all one-liners from the directory "one-liners"
	oneLiners, err := LoadOneLiners("one-liners")
	if err != nil {
		fmt.Printf("Error loading one-liners: %v\n", err)
		return
	}

	for {
		fmt.Println("Available One-Liners:")
		fmt.Println("----------------------")
		for _, ol := range oneLiners {
			fmt.Printf("Name: %s\nDescription: %sID: %s\n\n", ol.Info.Name, ol.Info.Description, ol.ID)
		}

		prompt := promptui.Prompt{
			Label: "Enter a command (run <name_one_liner> or exit)",
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if strings.ToLower(result) == "exit" {
			break
		}

		if strings.HasPrefix(result, "run ") {
			oneLinerName := strings.TrimSpace(strings.TrimPrefix(result, "run "))
			found := false
			for _, ol := range oneLiners {
				if ol.Info.Name == oneLinerName {
					runOneLiner(ol.OneLiner.Cmd)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("One-liner '%s' not found.\n", oneLinerName)
			}
		} else {
			fmt.Println("Invalid command. Please use 'run <name_one_liner>' or 'exit'.")
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
