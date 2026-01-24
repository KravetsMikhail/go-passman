package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go-passman/internal/storage"
	"go-passman/internal/utils"
)

// NewEncryptCommand creates the encrypt command
func NewEncryptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleEncrypt()
		},
	}

	return cmd
}

// NewDecryptCommand creates the decrypt command
func NewDecryptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleDecrypt()
		},
	}

	return cmd
}

func handleEncrypt() error {
	vault, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if vault.Encrypted {
		fmt.Println("ℹ️  Vault is already encrypted.")
		return nil
	}

	// Get master password
	password, err := utils.ReadPasswordConfirm()
	if err != nil {
		return err
	}

	vault.Encrypted = true
	if err := storage.SaveVault(vault, &password); err != nil {
		return err
	}

	fmt.Println("✅ Vault encrypted successfully.")
	return nil
}

func handleDecrypt() error {
	isEncrypted, err := storage.IsVaultEncrypted()
	if err != nil {
		return err
	}

	if !isEncrypted {
		fmt.Println("ℹ️  Vault is not encrypted.")
		return nil
	}

	vault, err := storage.LoadVault()
	if err != nil {
		return err
	}

	vault.Encrypted = false
	if err := storage.SaveVault(vault, nil); err != nil {
		return err
	}

	fmt.Println("✅ Vault decrypted successfully.")
	return nil
}
