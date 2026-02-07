package cmd

import (
	"fmt"
	"os"

	"go-passman/internal/storage"

	"github.com/spf13/cobra"
)

// NewRootCommand creates the root command
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "go-passman",
		Short:   "A simple CLI password manager",
		Long:    "A simple and secure CLI password manager. Store, manage, encrypt, and decrypt passwords from your terminal.",
		Version: "0.2.0",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := storage.Init(); err != nil {
				fmt.Fprintf(os.Stderr, "Error initializing storage: %v\n", err)
				os.Exit(1)
			}
		},
	}

	// Add subcommands
	rootCmd.AddCommand(
		NewAddCommand(),
		NewRemoveCommand(),
		NewCopyCommand(),
		NewListCommand(),
		NewUpdateCommand(),
		NewOpenCommand(),
		NewPathCommand(),
		NewEncryptCommand(),
		NewDecryptCommand(),
		NewStatusCommand(),
	)

	return rootCmd
}
