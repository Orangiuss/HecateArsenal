package commands.install

// install is the command to install a new tool in the arsenal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Install installs a new tool in the arsenal
func Install(tool