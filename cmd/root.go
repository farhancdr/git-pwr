/*
Copyright © 2024 Farhan Sadek farhancdr@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-pwr",
	Short: "git helper with go",
	Long:  `A helper for your daily git usage`,
}

var listBranchesCount int
var switchBack bool

var copyBranchCmd = &cobra.Command{
	Use:   "copy-branch",
	Short: "List git branches",
	Run: func(cmd *cobra.Command, args []string) {
		execCopyBranch()
	},
}

var publishBranchCmd = &cobra.Command{
	Use:   "publish-branch",
	Short: "Create a new branch and publish to remote origin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter the branch name: ")
		reader := bufio.NewReader(os.Stdin)
		branchName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading branch name:", err)
			os.Exit(1)
		}
		branchName = strings.TrimSpace(branchName) // Remove trailing newline

		execPublishBranch(branchName)
	},
}

var deleteBranchCmd = &cobra.Command{
	Use:   "delete-all-branch",
	Short: "Delete all local branches except the current one",
	Run: func(cmd *cobra.Command, args []string) {
		execDeleteAllBranches()
	},
}

var pushTagCmd = &cobra.Command{
	Use:   "ptag [tag name]",
	Short: "Create and push a new tag to origin",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tagName := strings.TrimSpace(args[0])
		execPushTag(tagName)
	},
}

func runCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %v", err)
	}
	return string(output), nil
}

func Execute() {
	rootCmd.AddCommand(copyBranchCmd)
	rootCmd.AddCommand(publishBranchCmd)
	rootCmd.AddCommand(deleteBranchCmd)
	rootCmd.AddCommand(pushTagCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	copyBranchCmd.Flags().IntVarP(&listBranchesCount, "num", "n", 10, "Number of branches to show")
	publishBranchCmd.Flags().BoolVarP(&switchBack, "switch-back", "s", false, "Switch back to the previous branch after creating the new branch default=true")
}
