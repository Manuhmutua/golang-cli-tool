package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

// This command lists all the S3 buckets
var listS3Buckets = &cobra.Command{
	Use:   "s3",
	Short: "List all s3 buckets using this command",
	Long:  ` Use it like: cli-tool-golang list s3`,
	Run: func(cmd *cobra.Command, args []string) {

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_cmd := exec.Command(dir + "/cmd/bash/list_s3_buckets.sh")
		stdout, err := _cmd.Output()
		if err != nil {
			if err.Error() == "exit status 253" {
				fmt.Println("Please add a configuration file with valid contents first to continue")
				os.Exit(1)
			}
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf(string(stdout))
	},
}

func init() {
	listCmd.AddCommand(listS3Buckets)
}
