package cmd_install

// install is the command to install a new tool in the arsenal

import (
	"fmt"
	"os"
	"strings"
)

// Install installs a new tool in the arsenal, arsenalPath defaults to $HOME/.hecatearsenal
func toolInstall(toolName string, arsenalPath string) {
	// Check if the tool is already installed in path, with execution
	_, err := exec.LookPath(toolName)
	if err == nil {
		fmt.Printf("%s is already installed in the PATH\n", toolName)
		// Check if the tool is executable
		return
	}

	// Check if the tool is already installed in the arsenal
	toolPath := fmt.Sprintf("%s/%s", arsenalPath, toolName)
	if _, err := os.Stat(toolPath); err == nil {
		// Check if the tool is executable
		if _, err := exec.LookPath(toolPath); err == nil {
			fmt.Printf("%s is already installed\n", toolName)
			return
		}
		else {
			fmt.Printf("%s is already installed, but not executable\n", toolName)
			return
		}
	}

	// Check if the tool is in the PATH
	_, err = exec.LookPath(toolName)
	if err == nil {
		fmt.Printf("%s is already installed\n", toolName)
		return
	}
