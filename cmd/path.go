package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
)

// NewPathCommand creates the path command
func NewPathCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "path",
		Short: "Display the path to the vault file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handlePath()
		},
	}

	return cmd
}

func handlePath() error {
	path := storage.GetVaultPath()
	fmt.Printf("Vault path: %s\n", path)
	return nil
}
