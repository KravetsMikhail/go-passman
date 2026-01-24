package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
	"go-passman/internal/utils"
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
	vault, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if len(vault.Entries) == 0 {
		fmt.Println("📭 No services to update.")
		os.Exit(1)
	}

	// List services and let user choose
	services := make([]string, 0, len(vault.Entries))
	for service := range vault.Entries {
		services = append(services, service)
	}

	service, err := utils.ChooseFromList(services, "Select a service to update:")
	if err != nil {
		return err
	}

	entry := vault.Entries[service]

	// Ask if user wants to update login
	updateLogin, err := utils.ReadInput("Update login? (y/n, default n): ")
	if err != nil {
		return err
	}

	if updateLogin == "y" || updateLogin == "Y" {
		login, err := utils.ReadInput(fmt.Sprintf("Enter new login (current: %s): ", entry.Login))
		if err != nil {
			return err
		}
		entry.Login = login
	}

	password, err := utils.ReadPassword("Enter new password: ")
	if err != nil {
		return err
	}

	entry.Password = password
	vault.Entries[service] = entry

	if err := storage.SaveVault(vault, nil); err != nil {
		return err
	}

	fmt.Printf("✅ Password for '%s' updated.\n", service)
	return nil
}

func handleUpdateGenerate() error {
	vault, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if len(vault.Entries) == 0 {
		fmt.Println("📭 No services to update.")
		os.Exit(1)
	}

	// List services and let user choose
	services := make([]string, 0, len(vault.Entries))
	for service := range vault.Entries {
		services = append(services, service)
	}

	service, err := utils.ChooseFromList(services, "Select a service to update:")
	if err != nil {
		return err
	}

	entry := vault.Entries[service]

	// Ask if user wants to update login
	updateLogin, err := utils.ReadInput("Update login? (y/n, default n): ")
	if err != nil {
		return err
	}

	if updateLogin == "y" || updateLogin == "Y" {
		login, err := utils.ReadInput(fmt.Sprintf("Enter new login (current: %s): ", entry.Login))
		if err != nil {
			return err
		}
		entry.Login = login
	}

	// Get password generation options
	length, useNumbers, useSpecial := utils.ChoosePasswordOptions()
	password := utils.GeneratePassword(length, useNumbers, useSpecial)

	entry.Password = password
	vault.Entries[service] = entry

	if err := storage.SaveVault(vault, nil); err != nil {
		return err
	}

	// Copy to clipboard
	if err := utils.CopyToClipboard(password); err != nil {
		fmt.Printf("⚠️  Password updated but clipboard copy failed: %v\n", err)
	} else {
		fmt.Printf("✅ Password for '%s' updated and copied to clipboard.\n", service)
	}

	return nil
}
