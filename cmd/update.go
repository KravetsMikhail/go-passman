package cmd

import (
	"errors"
	"fmt"
	"time"

	"go-passman/internal/models"
	"go-passman/internal/storage"
	"go-passman/internal/utils"

	"github.com/spf13/cobra"
)

// NewUpdateCommand creates the update command
func NewUpdateCommand() *cobra.Command {
	var generate bool

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing service or entry in the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			if generate {
				return handleUpdateGenerate()
			}
			return handleUpdateManual()
		},
	}

	cmd.Flags().BoolVarP(&generate, "generate", "g", false, "Generate a new random password")

	return cmd
}

func handleUpdateManual() error {
	vault, pwd, err := storage.LoadVault()
	if err != nil {
		return err
	}

	for {
		if len(vault.Entries) == 0 {
			fmt.Println("📭 No services to update.")
			break
		}

		services := getSortedServices(vault.Entries)
		service, err := utils.ChooseFromList(services, "Select a service to update:", "Filter by name (Enter = all): ")
		if err != nil {
			if errors.Is(err, utils.ErrCancelled) {
				break
			}
			return err
		}

		entry := vault.Entries[service]

		// Login: show current; Enter = keep, type = replace
		login, err := utils.ReadInput(fmt.Sprintf("Login [%s]: ", entry.Login))
		if err != nil {
			return err
		}
		if login != "" {
			entry.Login = login
		}

		// Host: show current; Enter = keep, type = replace
		host, err := utils.ReadInput(fmt.Sprintf("Host [%s]: ", entry.Host))
		if err != nil {
			return err
		}
		if host != "" {
			entry.Host = host
		}

		// Comment: show current; Enter = keep, type = replace
		comment, err := utils.ReadInput(fmt.Sprintf("Comment [%s]: ", entry.Comment))
		if err != nil {
			return err
		}
		if comment != "" {
			entry.Comment = comment
		}

		// Password: Enter = keep current, type new = replace (current not shown)
		password, err := utils.ReadPassword("Password (Enter to keep current): ")
		if err != nil {
			return err
		}
		if password != "" {
			entry.Password = password
		}

		vault.Entries[service] = entry

		if err := storage.SaveVault(vault, pwd); err != nil {
			return err
		}

		fmt.Printf("✅ Password for '%s' updated.\n", service)
		printEntrySummary(service, &entry)

		printListCompact(getSortedServices(vault.Entries), vault.Entries)
		if !utils.ConfirmActionWithTimeout("Continue?", 30*time.Second) {
			break
		}
	}
	return nil
}

func handleUpdateGenerate() error {
	vault, pwd, err := storage.LoadVault()
	if err != nil {
		return err
	}

	for {
		if len(vault.Entries) == 0 {
			fmt.Println("📭 No services to update.")
			break
		}

		services := getSortedServices(vault.Entries)
		service, err := utils.ChooseFromList(services, "Select a service to update:", "Filter by name (Enter = all): ")
		if err != nil {
			if errors.Is(err, utils.ErrCancelled) {
				break
			}
			return err
		}

		entry := vault.Entries[service]

		// Login: show current; Enter = keep, type = replace
		login, err := utils.ReadInput(fmt.Sprintf("Login [%s]: ", entry.Login))
		if err != nil {
			return err
		}
		if login != "" {
			entry.Login = login
		}

		// Host: show current; Enter = keep, type = replace
		host, err := utils.ReadInput(fmt.Sprintf("Host [%s]: ", entry.Host))
		if err != nil {
			return err
		}
		if host != "" {
			entry.Host = host
		}

		// Comment: show current; Enter = keep, type = replace
		comment, err := utils.ReadInput(fmt.Sprintf("Comment [%s]: ", entry.Comment))
		if err != nil {
			return err
		}
		if comment != "" {
			entry.Comment = comment
		}

		// Generate new password (replaces current)
		length, useNumbers, useSpecial := utils.ChoosePasswordOptions()
		password := utils.GeneratePassword(length, useNumbers, useSpecial)
		entry.Password = password

		vault.Entries[service] = entry

		if err := storage.SaveVault(vault, pwd); err != nil {
			return err
		}

		// Copy to clipboard
		if err := utils.CopyToClipboard(password); err != nil {
			fmt.Printf("⚠️ Password updated but clipboard copy failed: %v\n", err)
		} else {
			fmt.Printf("✅ Password for '%s' updated and copied to clipboard.\n", service)
		}
		printEntrySummary(service, &entry)

		printListCompact(getSortedServices(vault.Entries), vault.Entries)
		if !utils.ConfirmActionWithTimeout("Continue?", 30*time.Second) {
			break
		}
	}
	return nil
}

// printEntrySummary prints the entry fields after update (password masked).
func printEntrySummary(service string, entry *models.PasswordEntry) {
	fmt.Println()
	fmt.Println("  New values:")
	fmt.Printf("    Service: %s\n", service)
	fmt.Printf("    Login:   %s\n", orEmpty(entry.Login))
	fmt.Printf("    Host:    %s\n", orEmpty(entry.Host))
	fmt.Printf("    Comment: %s\n", orEmpty(entry.Comment))
	fmt.Println("    Password: ****")
	fmt.Println()
}

func orEmpty(s string) string {
	if s == "" {
		return "(empty)"
	}
	return s
}
