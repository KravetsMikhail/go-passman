package cmd

import (
	"fmt"
	"os"

	"go-passman/internal/models"
	"go-passman/internal/storage"
	"go-passman/internal/utils"

	"github.com/spf13/cobra"
)

// NewAddCommand creates the add command
func NewAddCommand() *cobra.Command {
	var generate bool

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new service or entry to the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			if generate {
				return handleAddGenerate()
			}
			return handleAddManual()
		},
	}

	cmd.Flags().BoolVarP(&generate, "generate", "g", false, "Generate a random password")

	return cmd
}

func handleAddManual() error {
	service, err := utils.ReadInput("Enter service name: ")
	if err != nil {
		return err
	}

	// Check if service already exists
	vault, pwd, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if _, exists := vault.Entries[service]; exists {
		fmt.Printf("❌ Service '%s' already exists. Use 'update' to change it.\n", service)
		os.Exit(1)
	}

	login, err := utils.ReadInput("Enter login (optional, press Enter to skip): ")
	if err != nil {
		return err
	}

	host, err := utils.ReadInput("Enter host (optional, press Enter to skip): ")
	if err != nil {
		return err
	}

	comment, err := utils.ReadInput("Enter comment (optional, press Enter to skip): ")
	if err != nil {
		return err
	}

	password, err := utils.ReadPassword("Enter password: ")
	if err != nil {
		return err
	}

	vault.Entries[service] = models.PasswordEntry{
		Login:     login,
		Host:      host,
		Comment:   comment,
		Password:  password,
		Encrypted: false,
	}

	if err := storage.SaveVault(vault, pwd); err != nil {
		return err
	}

	fmt.Printf("✅ Password for '%s' saved.\n", service)
	return nil
}

func handleAddGenerate() error {
	service, err := utils.ReadInput("Enter service name: ")
	if err != nil {
		return err
	}

	// Check if service already exists
	vault, pwd, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if _, exists := vault.Entries[service]; exists {
		fmt.Printf("❌ Service '%s' already exists. Use 'update' to change it.\n", service)
		os.Exit(1)
	}

	login, err := utils.ReadInput("Enter login (optional, press Enter to skip): ")
	if err != nil {
		return err
	}

	host, err := utils.ReadInput("Enter host (optional, press Enter to skip): ")
	if err != nil {
		return err
	}

	comment, err := utils.ReadInput("Enter comment (optional, press Enter to skip): ")
	if err != nil {
		return err
	}

	// Get password generation options
	length, useNumbers, useSpecial := utils.ChoosePasswordOptions()
	password := utils.GeneratePassword(length, useNumbers, useSpecial)

	vault.Entries[service] = models.PasswordEntry{
		Login:     login,
		Host:      host,
		Comment:   comment,
		Password:  password,
		Encrypted: false,
	}

	if err := storage.SaveVault(vault, pwd); err != nil {
		return err
	}

	// Copy to clipboard
	if err := utils.CopyToClipboard(password); err != nil {
		fmt.Printf("⚠️  Password saved but clipboard copy failed: %v\n", err)
	} else {
		fmt.Printf("✅ Password for '%s' saved and copied to clipboard.\n", service)
	}

	return nil
}
