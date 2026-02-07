package models

// PasswordEntry represents a single password entry in the vault
type PasswordEntry struct {
	Login     string `json:"login,omitempty"`
	Host      string `json:"host,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Password  string `json:"password"`
	Encrypted bool   `json:"encrypted"`
}

// Vault represents the entire password vault
type Vault struct {
	Entries   map[string]PasswordEntry `json:"entries"`
	Encrypted bool                     `json:"encrypted"`
}

// NewVault creates a new empty vault
func NewVault() *Vault {
	return &Vault{
		Entries:   make(map[string]PasswordEntry),
		Encrypted: false,
	}
}
