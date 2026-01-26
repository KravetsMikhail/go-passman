package storage

import (
	"encoding/json"
	"fmt"
	"go-passman/internal/crypto"
	"go-passman/internal/models"
	"os"
	"os/exec"
	"path/filepath"
)

var vaultPath string

// Init initializes the vault path (same directory as executable)
func Init() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	exeDir := filepath.Dir(exePath)
	vaultPath = filepath.Join(exeDir, "vault.json")
	return nil
}

// GetVaultPath returns the path to the vault file
func GetVaultPath() string {
	return vaultPath
}

// LoadVault loads the vault from disk, decrypting if necessary
func LoadVault() (*models.Vault, error) {
	if _, err := os.Stat(vaultPath); os.IsNotExist(err) {
		return models.NewVault(), nil
	}

	data, err := os.ReadFile(vaultPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read vault: %w", err)
	}

	// Try to parse as JSON first
	var vault models.Vault
	if err := json.Unmarshal(data, &vault); err == nil {
		if vault.Encrypted {
			password, err := readPassword("Vault is encrypted. Please enter your password: ")
			if err != nil {
				return nil, fmt.Errorf("failed to read password: %w", err)
			}

			decrypted, err := crypto.Decrypt(password, string(data))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error decrypting vault: %v\n", err)
				os.Exit(1)
			}

			if err := json.Unmarshal(decrypted, &vault); err != nil {
				return nil, fmt.Errorf("failed to parse decrypted vault: %w", err)
			}
			return &vault, nil
		}
		return &vault, nil
	}

	// If JSON parsing fails, assume it's encrypted
	password, err := readPassword("Vault seems encrypted or corrupted. Please enter password: ")
	if err != nil {
		return nil, fmt.Errorf("failed to read password: %w", err)
	}

	decrypted, err := crypto.Decrypt(password, string(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decrypting vault: %v\n", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(decrypted, &vault); err != nil {
		return nil, fmt.Errorf("failed to parse decrypted vault: %w", err)
	}

	return &vault, nil
}

// SaveVault saves the vault to disk, encrypting if necessary
func SaveVault(vault *models.Vault, password *string) error {
	vaultJSON, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}

	if vault.Encrypted {
		if password == nil {
			return fmt.Errorf("password required for encrypted vault")
		}
		encrypted, err := crypto.Encrypt(*password, vaultJSON)
		if err != nil {
			return fmt.Errorf("encryption error: %w", err)
		}
		if err := os.WriteFile(vaultPath, []byte(encrypted), 0600); err != nil {
			return fmt.Errorf("failed to write encrypted vault: %w", err)
		}
	} else {
		if err := os.WriteFile(vaultPath, vaultJSON, 0600); err != nil {
			return fmt.Errorf("failed to write vault: %w", err)
		}
	}

	return nil
}

// OpenInEditor opens the vault in the specified editor
func OpenInEditor(editor string) error {
	vault, err := LoadVault()
	if err != nil {
		return err
	}

	vaultJSON, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}

	// Write decrypted content to temp file
	tempFile, err := os.CreateTemp("", "vault-*.json")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(vaultJSON); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}
	tempFile.Close()

	// Open in editor
	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to open editor: %w", err)
	}

	// Read back the modified content
	modifiedData, err := os.ReadFile(tempFile.Name())
	if err != nil {
		return fmt.Errorf("failed to read modified vault: %w", err)
	}

	// Update vault from modified data
	if err := json.Unmarshal(modifiedData, vault); err != nil {
		return fmt.Errorf("invalid JSON in vault file: %w", err)
	}

	// Save the vault
	return SaveVault(vault, nil)
}

// readPassword reads a password from stdin securely
func readPassword(prompt string) (string, error) {
	fmt.Print(prompt)

	// Use stty to hide password input on Unix-like systems
	cmd := exec.Command("stty", "-echo")
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		// Fallback for Windows or if stty is not available
		// Just read normally (not ideal but will work)
		return readPasswordFallback()
	}
	//defer exec.Command("stty", "echo").Run()
	defer func() {
		cmd = exec.Command("stty", "echo")
		cmd.Stdin = os.Stdin
		cmd.Run()
	}()

	var password string
	fmt.Scanln(&password)
	fmt.Println() // New line after password

	return password, nil
}

// readPasswordFallback reads password without hiding it (fallback for Windows)
func readPasswordFallback() (string, error) {
	var password string
	fmt.Scanln(&password)
	return password, nil
}

// IsVaultEncrypted checks if the vault is encrypted
func IsVaultEncrypted() (bool, error) {
	if _, err := os.Stat(vaultPath); os.IsNotExist(err) {
		return false, nil
	}

	data, err := os.ReadFile(vaultPath)
	if err != nil {
		return false, fmt.Errorf("failed to read vault: %w", err)
	}

	var vault models.Vault
	if err := json.Unmarshal(data, &vault); err == nil {
		return vault.Encrypted, nil
	}

	return true, nil
}
