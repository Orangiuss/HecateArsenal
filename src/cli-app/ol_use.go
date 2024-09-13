package main_cli

import (
	"fmt"

	one_liners "github.com/Orangiuss/hecate/src/core/one-liners"

	"github.com/spf13/cobra"
)

var useOneLinerCmd = &cobra.Command{
	Use:   "ol use <one-liner_name>",
	Short: "Execute a one-liner",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		oneLinerName := args[0]
		oneLiners, err := one_liners.LoadOneLiners("one-liners-templates")
		if err != nil {
			fmt.Println("Error reading one-liners:", err)
			return
		}
		oneLiner, err := one_liners.RunOneLiner(oneLiners, oneLinerName, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Execute the one-liner (this needs to be implemented)
		fmt.Printf("Executing one-liner: %s\n", oneLiner.Name)
	},
}
