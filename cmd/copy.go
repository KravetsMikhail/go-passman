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
		Use:   "copy [service|N]",
		Short: "Copy the password to the clipboard (by name or by number from list)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleCopy(args[0])
		},
	}

	return cmd
}

func handleCopy(serviceOrNum string) error {
	vault, _, err := storage.LoadVault()
	if err != nil {
		return err
	}

	service, err := resolveServiceOrNumber(vault.Entries, serviceOrNum)
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		os.Exit(1)
	}

	entry := vault.Entries[service]

	if err := utils.CopyToClipboard(entry.Password); err != nil {
		return err
	}

	if entry.Login != "" {
		fmt.Printf("Login for '%s': %s\n", service, entry.Login)
	}
	if entry.Host != "" {
		fmt.Printf("Host for '%s': %s\n", service, entry.Host)
	}
	if entry.Comment != "" {
		fmt.Printf("Comment for '%s': %s\n", service, entry.Comment)
	}

	fmt.Printf("📋 Password for '%s' copied to clipboard!\n", service)

	return nil
}
