package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// ReadInput reads a line of input from stdin
func ReadInput(prompt string) (string, error) {
	fmt.Print(prompt)
	//var input string
	//_, err := fmt.Scanln(&input)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // Reads entire line including spaces
	input = strings.TrimSpace(input)      // Trim newline and spaces
	return input, err
}

// ReadPassword reads a password without echoing (hidden input).
// Works on Windows, Linux, and macOS when stdin is a terminal.
// Falls back to visible input if terminal is not available (e.g. redirected stdin).
func ReadPassword(prompt string) (string, error) {
	fmt.Print(prompt)

	fd := int(os.Stdin.Fd())
	if term.IsTerminal(fd) {
		bytePassword, err := term.ReadPassword(fd)
		fmt.Println() // newline after hidden input
		if err != nil {
			return "", err
		}
		return string(bytePassword), nil
	}

	// Fallback when stdin is not a terminal (e.g. pipe, redirect)
	var password string
	_, err := fmt.Scanln(&password)
	if err != nil {
		return "", err
	}
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
