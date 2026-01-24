package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
)

// NewStatusCommand creates the status command
func NewStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Display the status of the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleStatus()
		},
	}

	return cmd
}

func handleStatus() error {
	vault, err := storage.LoadVault()
	if err != nil {
		return err
	}

	fmt.Println("🔐 Vault Status:")
	fmt.Printf("  Entries: %d\n", len(vault.Entries))
	fmt.Printf("  Encrypted: %v\n", vault.Encrypted)
	fmt.Printf("  Path: %s\n", storage.GetVaultPath())

	return nil
}
