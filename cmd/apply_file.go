package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply AWS config files",
	Long: `Use this command to set the aws configurations before running the list command. For example:
cli-tool-golang apply -f <Absolute Path of the configuration file>.`,
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

	if len(args) == 1 {

		_ = os.Setenv("FILE", args[0])

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		cmd := exec.Command(dir + "/cmd/check_file.sh")
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

		cmd = exec.Command(dir + "/cmd/aws_configure.sh")
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
