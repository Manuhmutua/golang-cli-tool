package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// addCmd represents the add command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply AWS config files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fstatus, _ := cmd.Flags().GetBool("file")

		if fstatus {
			addFile(args)
		} else {
			fmt.Println("Command must contain -f or --file flag")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().BoolP("file", "f", false, "Add config file")
}

func addFile(args []string) {

	if len(args) > 1 {
		fmt.Println("Only one argument is allowed - The file path")
		os.Exit(1)
	}

	fmt.Printf("File location is: %s \n", args)
}