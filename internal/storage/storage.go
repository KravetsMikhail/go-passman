package storage

import (
	"encoding/json"
	"fmt"
	"go-passman/internal/crypto"
	"go-passman/internal/models"
	"go-passman/internal/utils"
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

// LoadVault loads the vault from disk, decrypting if necessary (prompts for password on terminal).
// When the vault is encrypted, the password used for decryption is returned as second value
// so callers can pass it to SaveVault when saving (avoids asking for password twice).
func LoadVault() (*models.Vault, *string, error) {
	return loadVaultWithPassword(nil, true)
}

// LoadVaultWithPassword loads the vault using the given password when encrypted.
// If vault is encrypted and password is nil, returns (nil, nil, err) so the caller can ask for password (e.g. web unlock form).
// If vault is encrypted and password is provided, uses it and returns (vault, &password, nil).
func LoadVaultWithPassword(password *string) (*models.Vault, *string, error) {
	return loadVaultWithPassword(password, false)
}

func loadVaultWithPassword(password *string, promptIfEncrypted bool) (*models.Vault, *string, error) {
	if _, err := os.Stat(vaultPath); os.IsNotExist(err) {
		return models.NewVault(), nil, nil
	}

	data, err := os.ReadFile(vaultPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read vault: %w", err)
	}

	var vault models.Vault
	if err := json.Unmarshal(data, &vault); err == nil {
		if !vault.Encrypted {
			return &vault, nil, nil
		}
		// Encrypted JSON metadata read; body is still encrypted - try decrypt with provided or prompt
		var pwd string
		if password != nil {
			pwd = *password
		} else if promptIfEncrypted {
			var errPwd error
			pwd, errPwd = utils.ReadPassword("Vault is encrypted. Please enter your password: ")
			if errPwd != nil {
				return nil, nil, fmt.Errorf("failed to read password: %w", errPwd)
			}
		} else {
			return nil, nil, fmt.Errorf("vault is encrypted: password required")
		}
		decrypted, err := crypto.Decrypt(pwd, string(data))
		if err != nil {
			if promptIfEncrypted {
				fmt.Fprintf(os.Stderr, "Error decrypting vault: %v\n", err)
				os.Exit(1)
			}
			return nil, nil, fmt.Errorf("decryption failed: %w", err)
		}
		if err := json.Unmarshal(decrypted, &vault); err != nil {
			return nil, nil, fmt.Errorf("failed to parse decrypted vault: %w", err)
		}
		return &vault, &pwd, nil
	}

	// Invalid JSON — assume encrypted blob
	if password != nil {
		decrypted, err := crypto.Decrypt(*password, string(data))
		if err != nil {
			return nil, nil, fmt.Errorf("decryption failed: %w", err)
		}
		if err := json.Unmarshal(decrypted, &vault); err != nil {
			return nil, nil, fmt.Errorf("failed to parse decrypted vault: %w", err)
		}
		return &vault, password, nil
	}
	if promptIfEncrypted {
		pwd, errPwd := utils.ReadPassword("Vault seems encrypted or corrupted. Please enter password: ")
		if errPwd != nil {
			return nil, nil, fmt.Errorf("failed to read password: %w", errPwd)
		}
		decrypted, err := crypto.Decrypt(pwd, string(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decrypting vault: %v\n", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(decrypted, &vault); err != nil {
			return nil, nil, fmt.Errorf("failed to parse decrypted vault: %w", err)
		}
		return &vault, &pwd, nil
	}
	return nil, nil, fmt.Errorf("vault is encrypted: password required")
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
	vault, pwd, err := LoadVault()
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

	// Save the vault (pass password if vault was encrypted)
	return SaveVault(vault, pwd)
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
