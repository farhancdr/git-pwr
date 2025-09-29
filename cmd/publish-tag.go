package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execPushTag(tagName string) {
	tagName = strings.TrimSpace(tagName)
	if tagName == "" {
		fmt.Printf("%sError: tag name cannot be empty.%s\n", red, reset)
		return
	}

	// Create the tag
	fmt.Printf("%sCreating tag '%s'...%s\n", yellow, tagName, reset)
	createCmd := exec.Command("git", "tag", tagName)
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr
	if err := createCmd.Run(); err != nil {
		fmt.Printf("%sFailed to create tag '%s': %v%s\n", red, tagName, err, reset)
		return
	}

	// Push the tag to origin
	fmt.Printf("%sPushing tag '%s' to origin...%s\n", yellow, tagName, reset)
	pushCmd := exec.Command("git", "push", "origin", tagName)
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		fmt.Printf("%sFailed to push tag '%s' to origin: %v%s\n", red, tagName, err, reset)
		return
	}

	fmt.Printf("%sTag '%s' successfully created and pushed to origin.%s\n", green, tagName, reset)
}
