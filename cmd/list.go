package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
)

// NewListCommand creates the list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all services or entries in the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleList()
		},
	}

	return cmd
}

func handleList() error {
	vault, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if len(vault.Entries) == 0 {
		fmt.Println("📭 No passwords saved yet.")
	} else {
		fmt.Println("🔐 Saved services:")
		for service, entry := range vault.Entries {
			if entry.Login != "" {
				fmt.Printf("- %s (login: %s)\n", service, entry.Login)
			} else {
				fmt.Printf("- %s\n", service)
			}
		}
	}

	return nil
}
