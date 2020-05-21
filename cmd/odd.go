package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// oddCmd represents the odd command
var oddCmd = &cobra.Command{
	Use:   "odd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var oddSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 != 0 {
				oddSum = oddSum + itemp
			}
		}

		fmt.Printf("The odd addition of %s is %d \n", args, oddSum)
	},
}

func init() {
	addCmd.AddCommand(oddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
