package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execPublishBranch(branchName string) {
	if branchName == "" {
		fmt.Println("Branch name is required")
		return
	}

	// Get the current branch
	currentBranch := getCurrentBranch()

	// Check if branch exists
	branchExists := isBranchExists(branchName)
	if !branchExists {
		createBranchCommand := exec.Command("git", "checkout", "-b", branchName)
		createBranchCommand.Stdout = os.Stdout
		createBranchCommand.Stderr = os.Stderr
		if err := createBranchCommand.Run(); err != nil {
			fmt.Println("Error creating branch:", err)
			return
		}
		fmt.Println("Branch created:", branchName)
	} else {
		fmt.Println("Branch already exists:", branchName)
	}

	// Get the remote URL from the existing configuration
	remoteURL := getRemoteURL("origin")

	// Check if remote exists
	remoteExists := isRemoteExists("origin")

	if !remoteExists {
		addRemoteCommand := exec.Command("git", "remote", "add", "origin", remoteURL)
		addRemoteCommand.Stdout = os.Stdout
		addRemoteCommand.Stderr = os.Stderr
		if err := addRemoteCommand.Run(); err != nil {
			fmt.Println("Error adding remote:", err)
			return
		}
		fmt.Println("Remote 'origin' added")
	}

	// Push branch to remote
	pushCommand := exec.Command("git", "push", "-u", "origin", branchName)
	pushCommand.Stdout = os.Stdout
	pushCommand.Stderr = os.Stderr
	if err := pushCommand.Run(); err != nil {
		fmt.Println("Error pushing branch to remote:", err)
		return
	}

	fmt.Printf("Branch '%s' created and pushed to remote 'origin'\n", branchName)

	// Switch back to the previous branch if requested
	if switchBack && currentBranch != "" {
		switchBackCommand := exec.Command("git", "checkout", currentBranch)
		switchBackCommand.Stdout = os.Stdout
		switchBackCommand.Stderr = os.Stderr
		if err := switchBackCommand.Run(); err != nil {
			fmt.Println("Error switching back to the previous branch:", err)
			return
		}
		fmt.Printf("Switched back to the previous branch: %s\n", currentBranch)
	}
}

func getCurrentBranch() string {
	currentBranchCommand := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := currentBranchCommand.Output()
	if err != nil {
		fmt.Println("Error getting current branch:", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

func isBranchExists(branchName string) bool {
	checkBranchCommand := exec.Command("git", "rev-parse", "--verify", branchName)
	checkBranchCommand.Stdout = os.Stdout
	checkBranchCommand.Stderr = os.Stderr
	err := checkBranchCommand.Run()
	return err == nil
}

func isRemoteExists(remoteName string) bool {
	listRemotesCommand := exec.Command("git", "remote")
	output, err := listRemotesCommand.Output()
	if err != nil {
		fmt.Println("Error listing remotes:", err)
		return false
	}
	remotes := strings.Fields(string(output))
	for _, remote := range remotes {
		if remote == remoteName {
			return true
		}
	}
	return false
}

func getRemoteURL(remoteName string) string {
	remoteURLCommand := exec.Command("git", "config", "--get", fmt.Sprintf("remote.%s.url", remoteName))
	output, err := remoteURLCommand.Output()
	if err != nil {
		fmt.Println("Error getting remote URL:", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}
