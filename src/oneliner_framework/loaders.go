package one_liners

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadOneLiners(directory string) ([]OneLiner, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var oneLiners []OneLiner
	for _, file := range files {
		content, err := ioutil.ReadFile(directory + "/" + file.Name())
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", file.Name(), err)
			continue
		}

		var oneLiner OneLiner
		err = yaml.Unmarshal(content, &oneLiner)
		if err != nil {
			fmt.Printf("Failed to parse YAML in file %s: %v\n", file.Name(), err)
			continue
		}

		oneLiner.ID = file.Name()
		oneLiners = append(oneLiners, oneLiner)
	}

	return oneLiners, nil
}
