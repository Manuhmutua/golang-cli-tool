package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// This is the apply file command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply AWS config files",
	Long: `Use this command to set the aws configurations before running the list command. For example:
cli-tool-golang apply -f <Absolute Path of the configuration file>.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Here we check if the user has used the --file/ -f command
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
	// here we set the flag to accept the configuration file path
	applyCmd.Flags().BoolP("file", "f", false, "Add config file")
}

// Here we validate file location and set file contents as the aws configure credentials
func addFile(args []string) {

	// Here we restrict using only one argument before/after the -f/--file flag
	if len(args) == 1 {

		_ = os.Setenv("FILE", args[0])

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Here we check if a file exists in the path provided by user
		cmd := exec.Command(dir + "/cmd/bash/check_file.sh")
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if string(stdout) == "does not exist\n" {
			fmt.Println("The file does not exist - please enter a valid file path")
			os.Exit(1)
		}

		cmd = exec.Command("cat", args[0])
		stdout, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Here we use the contents of the file to configure aws CLI
		for _, element := range strings.Split(string(stdout), "\n") {
			var variable = strings.Replace(strings.Split(element, "=")[0], " ", "", -1)
			var value = strings.Replace(strings.Split(element, "=")[1], " ", "", -1)
			if variable == "AWS_ACCESS_KEY_ID" {
				_ = os.Setenv("AWS_ACCESS_KEY_ID", value)
			} else if variable == "AWS_SECRET_ACCESS_KEY" {
				_ = os.Setenv("AWS_SECRET_ACCESS_KEY", value)
			} else if variable == "AWS_REGION" {
				_ = os.Setenv("AWS_REGION", value)
			}
		}

		cmd = exec.Command(dir + "/cmd/bash/aws_configure.sh")
		stdout, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(string(stdout))
		os.Exit(1)

	}

	fmt.Println("Only one argument is allowed - The file path")
	os.Exit(1)
}
