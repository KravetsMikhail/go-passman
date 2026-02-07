package cmd

import (
	"fmt"
	"os"

	"go-passman/internal/storage"
	"go-passman/internal/utils"

	"github.com/spf13/cobra"
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
	vault, _, err := storage.LoadVault()
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
		fmt.Printf("Login for '%s':		%s\n", service, entry.Login)
	}
	if entry.Host != "" {
		fmt.Printf("Host for '%s':		%s\n", service, entry.Host)
	}
	if entry.Comment != "" {
		fmt.Printf("Comment for '%s':		%s\n", service, entry.Comment)
	}

	fmt.Printf("📋 Password for '%s' copied to clipboard!\n", service)

	return nil
}
