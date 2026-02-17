package cmd

import (
	"fmt"
	"os"

	"go-passman/internal/storage"
	"go-passman/internal/web"

	"github.com/spf13/cobra"
)

// NewRootCommand creates the root command
func NewRootCommand() *cobra.Command {
	var runWeb bool
	rootCmd := &cobra.Command{
		Use:     "go-passman",
		Short:   "A simple CLI password manager",
		Long:    "A simple and secure CLI password manager. Store, manage, encrypt, and decrypt passwords from your terminal.",
		Version: "0.3.1",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if runWeb {
				return
			}
			if err := storage.Init(); err != nil {
				fmt.Fprintf(os.Stderr, "Error initializing storage: %v\n", err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if runWeb {
				if err := storage.Init(); err != nil {
					return fmt.Errorf("initializing storage: %w", err)
				}
				web.Run()
			}
			return nil
		},
	}

	rootCmd.Flags().BoolVarP(&runWeb, "web", "w", false, "Run as web server (simple UI at http://127.0.0.1:8080)")

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
