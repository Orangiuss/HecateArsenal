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

// runOneLiner executes the command associated with the one-liner, and store the result in a json file
func runOneLinerJson(cmd string, id string) {
	fmt.Printf("Running: %s\n", cmd)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}
	fmt.Printf("Output:\n%s\n", string(output))
	// Store the result in a json file
	storeResultInJsonFile(id, string(output))
}

// storeResultInJsonFile stores the result in a json file
func storeResultInJsonFile(id string, result string) {
	// Create a file
	file, err := os.Create(fmt.Sprintf("results/%s.json", id))
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write the result in the file
	_, err = file.WriteString(result)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}