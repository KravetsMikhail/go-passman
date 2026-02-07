package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
	"go-passman/internal/utils"
)

// NewRemoveCommand creates the remove command
func NewRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [service]",
		Short: "Remove a service or entry from the vault",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			service := args[0]
			return handleRemove(service)
		},
	}

	return cmd
}

func handleRemove(service string) error {
	vault, pwd, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if _, exists := vault.Entries[service]; !exists {
		fmt.Printf("❌ Service '%s' not found.\n", service)
		os.Exit(1)
	}

	if !utils.ConfirmAction(fmt.Sprintf("Are you sure you want to remove '%s'?", service)) {
		fmt.Println("❌ Operation cancelled.")
		return nil
	}

	delete(vault.Entries, service)

	if err := storage.SaveVault(vault, pwd); err != nil {
		return err
	}

	fmt.Printf("✅ Service '%s' removed.\n", service)
	return nil
}
