package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
	"go-passman/internal/utils"
)

// NewCopyCommand creates the copy command
func NewCopyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy [service]",
		Short: "Copy the password of a service to the clipboard",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			service := args[0]
			return handleCopy(service)
		},
	}

	return cmd
}

func handleCopy(service string) error {
	vault, err := storage.LoadVault()
	if err != nil {
		return err
	}

	entry, exists := vault.Entries[service]
	if !exists {
		fmt.Printf("❌ Service '%s' not found.\n", service)
		os.Exit(1)
	}

	if err := utils.CopyToClipboard(entry.Password); err != nil {
		return err
	}

	if entry.Login != "" {
		fmt.Printf("📋 Password for '%s' copied to clipboard! (login: %s)\n", service, entry.Login)
	} else {
		fmt.Printf("📋 Password for '%s' copied to clipboard!\n", service)
	}
	return nil
}
