package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ReadInput reads a line of input from stdin
func ReadInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}

// ReadPassword reads a password securely
func ReadPassword(prompt string) (string, error) {
	fmt.Print(prompt)

	// Try to hide password input using stty (Unix-like systems)
	cmd := exec.Command("stty", "-echo")
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		// Fallback for Windows: just read normally
		var password string
		fmt.Scanln(&password)
		return password, nil
	}
	defer exec.Command("stty", "echo").Run()

	var password string
	fmt.Scanln(&password)
	fmt.Println()

	return password, nil
}

// ReadPasswordConfirm reads a password twice and confirms they match
func ReadPasswordConfirm() (string, error) {
	password1, err := ReadPassword("Enter master password: ")
	if err != nil {
		return "", err
	}

	password2, err := ReadPassword("Confirm password: ")
	if err != nil {
		return "", err
	}

	if password1 != password2 {
		return "", fmt.Errorf("passwords do not match")
	}

	return password1, nil
}

// ChooseFromList lets user choose from a list of options
func ChooseFromList(items []string, prompt string) (string, error) {
	if len(items) == 0 {
		return "", fmt.Errorf("no items to choose from")
	}

	fmt.Println(prompt)
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Print("Enter your choice (number): ")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return "", err
	}

	if choice < 1 || choice > len(items) {
		return "", fmt.Errorf("invalid choice")
	}

	return items[choice-1], nil
}

// ConfirmAction asks user for confirmation
func ConfirmAction(message string) bool {
	fmt.Print(message + " (y/n): ")
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}
