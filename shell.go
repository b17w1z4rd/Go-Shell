package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display prompt
		fmt.Print("> ")

		// Read user input
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		// Exit shell if the user types "exit"
		if input == "exit" {
			break
		}

		// Split input into command and arguments
		cmdParts := splitCommand(input)
		if len(cmdParts) == 0 {
			continue
		}

		// Execute command
		cmd := exec.Command(cmdParts[0], cmdParts[1:]...) // first part is the command, rest are arguments
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
	}
}

// splitCommand splits the input string into command and arguments.
func splitCommand(input string) []string {
	var cmdParts []string
	var currentPart string

	for _, char := range input {
		if char == ' ' {
			if currentPart != "" {
				cmdParts = append(cmdParts, currentPart)
				currentPart = ""
			}
		} else {
			currentPart += string(char)
		}
	}
	if currentPart != "" {
		cmdParts = append(cmdParts, currentPart)
	}

	return cmdParts
}
