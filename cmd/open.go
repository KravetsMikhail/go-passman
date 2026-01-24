package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
)

// NewOpenCommand creates the open command
func NewOpenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open [editor]",
		Short: "Open the vault in a text editor",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			editor := "cat"
			if len(args) > 0 {
				editor = args[0]
			}
			return handleOpen(editor)
		},
	}

	return cmd
}

func handleOpen(editor string) error {
	if err := storage.OpenInEditor(editor); err != nil {
		return err
	}

	fmt.Println("✅ Vault updated.")
	return nil
}
