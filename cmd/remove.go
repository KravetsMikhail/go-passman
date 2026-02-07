package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
	"go-passman/internal/utils"
)

// NewRemoveCommand creates the remove command
func NewRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a service or entry (select from list)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleRemove()
		},
	}

	return cmd
}

func handleRemove() error {
	vault, pwd, err := storage.LoadVault()
	if err != nil {
		return err
	}

	for {
		if len(vault.Entries) == 0 {
			fmt.Println("📭 No services to remove.")
			break
		}

		services := getSortedServices(vault.Entries)
		service, err := utils.ChooseFromList(services, "Select a service to remove:", "Filter by name (Enter = all): ")
		if err != nil {
			if errors.Is(err, utils.ErrCancelled) {
				break
			}
			return err
		}

		if !utils.ConfirmAction(fmt.Sprintf("Are you sure you want to remove '%s'?", service)) {
			fmt.Println("❌ Operation cancelled.")
		} else {
			delete(vault.Entries, service)

			if err := storage.SaveVault(vault, pwd); err != nil {
				return err
			}

			fmt.Printf("✅ Service '%s' removed.\n", service)
		}

		printListCompact(getSortedServices(vault.Entries), vault.Entries)
		if !utils.ConfirmActionWithTimeout("Continue?", 30*time.Second) {
			break
		}
	}
	return nil
}
