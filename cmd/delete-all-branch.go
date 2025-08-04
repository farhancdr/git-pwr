package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	green  = "\033[0;32m"
	yellow = "\033[0;33m"
	red    = "\033[0;31m"
	reset  = "\033[0m"
)

func execDeleteAllBranches() {
	// Get current branch
	currentBranch, err := runCommand("git symbolic-ref --short HEAD")
	if err != nil {
		fmt.Printf("%sError: failed to get current branch: %v%s\n", red, err, reset)
		return
	}
	currentBranch = strings.TrimSpace(currentBranch)

	// Get all local branches
	allBranches, err := runCommand("git branch")
	if err != nil {
		fmt.Printf("%sError: failed to get local branches: %v%s\n", red, err, reset)
		return
	}

	var branchesToDelete []string
	scanner := bufio.NewScanner(strings.NewReader(allBranches))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		branchName := strings.TrimPrefix(line, "* ")
		if branchName != currentBranch && branchName != "" {
			branchesToDelete = append(branchesToDelete, branchName)
		}
	}

	if len(branchesToDelete) == 0 {
		fmt.Printf("%sNo branches to delete except '%s'.%s\n", green, currentBranch, reset)
		return
	}

	fmt.Printf("%sThe following branches will be deleted:%s\n", yellow, reset)
	for _, b := range branchesToDelete {
		fmt.Printf("%s- %s%s\n", yellow, b, reset)
	}

	fmt.Print("Are you sure you want to delete these branches? (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.ToLower(strings.TrimSpace(confirmation))

	if confirmation != "yes" {
		fmt.Printf("%sBranch deletion cancelled.%s\n", red, reset)
		return
	}

	for _, b := range branchesToDelete {
		cmd := exec.Command("git", "branch", "-D", b)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("%sFailed to delete branch %s: %v%s\n", red, b, err, reset)
		}
	}

	fmt.Printf("%sAll branches except '%s' have been deleted.%s\n", green, currentBranch, reset)
}
