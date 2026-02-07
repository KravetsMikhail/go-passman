package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"golang.org/x/term"
)

// ErrCancelled is returned when the user cancels an interactive choice (e.g. presses q).
var ErrCancelled = errors.New("cancelled")

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
// Restores terminal echo on Ctrl+C so the shell is not left in invisible input mode.
func ReadPassword(prompt string) (string, error) {
	fmt.Print(prompt)

	fd := int(os.Stdin.Fd())
	if term.IsTerminal(fd) {
		// Save terminal state before ReadPassword (which disables echo).
		// If user presses Ctrl+C, we restore it so the shell is usable again.
		oldState, err := term.GetState(fd)
		if err == nil {
			defer term.Restore(fd, oldState)
			// Restore terminal on SIGINT (Ctrl+C) so we don't leave echo disabled
			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt)
			defer signal.Stop(sigCh)
			go func() {
				<-sigCh
				term.Restore(fd, oldState)
				os.Exit(130) // 128 + 2 (SIGINT)
			}()
		}

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

const choosePageSize = 25 // max items shown at once when choosing; more than this triggers pagination

// ChooseFromList lets user choose from a list. If filterPrompt is non-empty, asks for a name filter first (substring, Enter = all).
// When there are more than choosePageSize items, shows pages (n = next, q = quit).
func ChooseFromList(items []string, prompt, filterPrompt string) (string, error) {
	if len(items) == 0 {
		return "", fmt.Errorf("no items to choose from")
	}

	list := items
	if filterPrompt != "" {
		filterLine, err := ReadInput(filterPrompt)
		if err != nil {
			return "", err
		}
		filter := strings.TrimSpace(strings.ToLower(filterLine))
		if filter != "" {
			filtered := make([]string, 0)
			for _, item := range items {
				if strings.Contains(strings.ToLower(item), filter) {
					filtered = append(filtered, item)
				}
			}
			if len(filtered) == 0 {
				fmt.Println("No matches. Showing all entries.")
			} else {
				list = filtered
			}
		}
	}

	if len(list) == 0 {
		return "", fmt.Errorf("no items to choose from")
	}

	if len(list) <= choosePageSize {
		fmt.Println(prompt)
		for i, item := range list {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Print("Enter your choice (number): ")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return "", err
		}
		if choice < 1 || choice > len(list) {
			return "", fmt.Errorf("invalid choice")
		}
		return list[choice-1], nil
	}

	// Paginated choice
	for page := 0; page*choosePageSize < len(list); page++ {
		start := page * choosePageSize
		end := start + choosePageSize
		if end > len(list) {
			end = len(list)
		}
		pageItems := list[start:end]
		totalPages := (len(list) + choosePageSize - 1) / choosePageSize
		fmt.Println(prompt)
		fmt.Printf("(page %d of %d, %d total)\n", page+1, totalPages, len(list))
		for i, item := range pageItems {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Printf("Enter choice (1-%d), n next page, q quit: ", len(pageItems))
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "q" {
			return "", ErrCancelled
		}
		if input == "n" {
			continue
		}
		var choice int
		if _, err := fmt.Sscanf(input, "%d", &choice); err != nil {
			fmt.Println("Invalid input.")
			page-- // re-show same page
			continue
		}
		if choice < 1 || choice > len(pageItems) {
			fmt.Println("Invalid choice.")
			page--
			continue
		}
		return pageItems[choice-1], nil
	}

	return "", ErrCancelled
}

// ConfirmAction asks user for confirmation
func ConfirmAction(message string) bool {
	fmt.Print(message + " (y/n): ")
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}

// ConfirmActionWithTimeout asks for confirmation with a timeout. If no answer within d, returns false
// and the caller should exit (e.g. so the vault is not left open in memory).
func ConfirmActionWithTimeout(message string, d time.Duration) bool {
	fmt.Print(message + " (y/n): ")
	resultCh := make(chan string, 1)
	go func() {
		var response string
		fmt.Scanln(&response)
		resultCh <- strings.TrimSpace(strings.ToLower(response))
	}()
	select {
	case response := <-resultCh:
		return response == "y"
	case <-time.After(d):
		fmt.Println()
		fmt.Printf("No response within %v. Exiting.\n", d)
		return false
	}
}
