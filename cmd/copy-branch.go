package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"
)

func listGitBranches(count int) ([]string, error) {
	localBranches, err := runCommand("git for-each-ref --sort=-committerdate --format '%(refname:short)' refs/heads | head -n " + fmt.Sprintf("%d", count))
	if err != nil {
		return nil, err
	}

	remoteBranches, err := runCommand("git for-each-ref --sort=-committerdate --format '%(refname:short)' refs/remotes | head -n " + fmt.Sprintf("%d", count))
	if err != nil {
		return nil, err
	}

	branches := append(strings.Split(localBranches, "\n"), strings.Split(remoteBranches, "\n")...)
	// Filter out any empty strings that may result from the splitting
	var filteredBranches []string
	for _, branch := range branches {
		if branch != "" {
			filteredBranches = append(filteredBranches, branch)
		}
	}

	return filteredBranches, nil
}

func printBranches(branches []string, selectedIndex int) {
	for i, branch := range branches {
		if i == selectedIndex {
			fmt.Printf("[x] \x1b[4;32m%s\x1b[0m\n", branch) // Green color for the selected branch
		} else {
			fmt.Printf("[ ] %s\n", branch)
		}
	}
}

func execCopyBranch() {
	keyboard.Open()
	defer keyboard.Close()

	branches, err := listGitBranches(listBranchesCount)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	selectedIndex := 0

	for {
		fmt.Print("\033[H\033[2J") // Clear the screen
		printBranches(branches, selectedIndex)
		// print the instructions
		fmt.Println("Use arrow keys to navigate, press enter to copy branch name to clipboard.")
		fmt.Println("Press ESC, Ctrl+C or q to exit.")

		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading keyboard input:", err)
			os.Exit(0)
			return
		}

		// if user press esc or ctrl+c, exit
		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC || char == 'q' {
			fmt.Println("Exiting...")
			os.Exit(0)
			return
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selectedIndex > 0 {
				selectedIndex--
			}
		case keyboard.KeyArrowDown:
			if selectedIndex < len(branches)-1 {
				selectedIndex++
			}
		case keyboard.KeyEnter:
			branchName := strings.TrimSpace(branches[selectedIndex])

			// Remove 'origin/' prefix if present
			branchName = strings.TrimPrefix(branchName, "origin/")

			err := clipboard.WriteAll(branchName)
			if err != nil {
				fmt.Println("Error copying to clipboard:", err)
			} else {
				fmt.Printf("\x1b[32mBranch '%s' copied to clipboard.\x1b[0m\n", branchName)
			}
			return
		}
	}
}
