package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources parent command",
	Long:  `The parent command for listing resources`,
	Run:   func(cmd *cobra.Command, args []string) {
		fmt.Printf("Specify the recource you want to list as an argument \n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
