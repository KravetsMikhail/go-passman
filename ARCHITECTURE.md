# Architecture Documentation

This document describes the architecture and design of go-passman.

## Overview

go-passman is a command-line password manager that securely stores and manages passwords using AES-256-GCM encryption. The application is written in Go and follows a modular, clean architecture.

## Directory Structure

```bash
go-passman/
├── main.go                    # Application entry point
├── cmd/                       # CLI command implementations
│   ├── root.go               # Root command and command registration
│   ├── add.go                # Add password command
│   ├── remove.go             # Remove password command
│   ├── copy.go               # Copy password to clipboard command
│   ├── list.go               # List passwords command
│   ├── update.go             # Update password command
│   ├── open.go               # Open vault in editor command
│   ├── encrypt.go            # Encrypt/decrypt vault commands
│   ├── status.go             # Show vault status command
│   └── path.go               # Show vault path command
├── internal/
│   ├── crypto/
│   │   └── crypto.go         # Encryption/decryption logic
│   ├── models/
│   │   └── models.go         # Data structures
│   ├── storage/
│   │   └── storage.go        # File I/O and vault management
│   └── utils/
│       ├── password.go       # Password generation
│       ├── clipboard.go      # Clipboard operations
│       └── interactive.go    # User interaction utilities
├── go.mod / go.sum           # Go module files
├── Makefile                  # Build automation
├── build.sh / build.bat      # Cross-platform build scripts
└── README.md                 # User documentation
```

## Key Components

### 1. CLI Framework (Cobra)

Uses [Cobra](https://github.com/spf13/cobra) for CLI argument parsing and command structure:

- Supports subcommands and flags
- Automatic help and version output
- Extensible command architecture

**Root Command** (`cmd/root.go`):

- Creates and registers all subcommands
- Initializes storage before any command execution

### 2. Data Models (`internal/models/models.go`)

```go
type PasswordEntry struct {
    Login     string `json:"login,omitempty"` //optional
    Host      string `json:"host,omitempty"` //optional
    Comment   string `json:"comment,omitempty"` //optional
    Password  string `json:"password"` // The actual password
    Encrypted bool   `json:"encrypted"` // Whether this entry is encrypted
}

type Vault struct {
    Entries  map[string]PasswordEntry // Service name -> Entry
    Encrypted bool                     // Whether entire vault is encrypted
}
```

### 3. Encryption/Decryption (`internal/crypto/crypto.go`)

**Algorithm**: AES-256-GCM

- **Encryption**: AES in Galois/Counter Mode (authenticated encryption)
- **Key Derivation**: PBKDF2-SHA256 with 100,000 iterations
- **Encoding**: Base64 for text storage

**Process**:

1. Generate random salt (16 bytes)
2. Derive key from password using PBKDF2
3. Generate random nonce (12 bytes)
4. Encrypt plaintext with AES-GCM
5. Combine salt + nonce + ciphertext
6. Encode as Base64

**Security Properties**:

- PBKDF2 iterations prevent brute-force attacks
- Random salt and nonce ensure identical plaintexts produce different ciphertexts
- AES-GCM provides both confidentiality and authenticity

### 4. Storage Management (`internal/storage/storage.go`)

**Vault Location**: `./vault.json` (same directory as executable)

**Functions**:

- `Init()` - Initialize vault path based on executable location
- `LoadVault()` - Load and decrypt vault from disk
- `SaveVault()` - Encrypt and save vault to disk
- `GetVaultPath()` - Get the vault file path
- `IsVaultEncrypted()` - Check encryption status
- `OpenInEditor()` - Open vault in text editor for manual editing

**Error Handling**:

- Attempts JSON parsing for unencrypted vaults
- Falls back to decryption attempt for encrypted vaults
- Returns descriptive errors for debugging

### 5. Command Implementations (`cmd/`)

Each command file implements a specific feature:

#### `add.go` - Add Password Entry

- Manual entry: Prompts for service name and password
- Generate: Prompts for options and generates secure password

#### `remove.go` - Remove Entry

- Requires confirmation before deletion
- Updates vault and saves to disk

#### `copy.go` - Copy to Clipboard

- Retrieves password from vault
- Copies to system clipboard

#### `list.go` - List All Entries

- Shows all service names in vault
- Displays count

#### `update.go` - Update Entry

- Lists available services
- Allows user to choose and update
- Supports generation or manual input

#### `open.go` - Open in Editor

- Decrypts vault
- Writes to temporary file
- Opens in specified editor
- Saves modified content back to vault

#### `encrypt.go` / `decrypt.go` - Toggle Encryption

- `encrypt`: Prompts for master password, encrypts vault
- `decrypt`: Removes encryption from vault

#### `status.go` - Show Status

- Displays vault statistics
- Shows encryption status
- Shows vault file path

#### `path.go` - Show Vault Path

- Outputs the path to vault.json

### 6. Utilities (`internal/utils/`)

#### `password.go` - Password Generation

- `GeneratePassword()` - Creates random passwords with options
- `ChoosePasswordOptions()` - Interactive prompt for generation settings

#### `clipboard.go` - Clipboard Operations

- Uses `atotto/clipboard` package
- Cross-platform support

#### `interactive.go` - User Interaction

- `ReadInput()` - Read simple text input
- `ReadPassword()` - Read password securely (hidden input)
- `ReadPasswordConfirm()` - Confirm password entry twice
- `ChooseFromList()` - Interactive selection from list
- `ConfirmAction()` - Yes/No confirmation

## Data Flow

### Adding a Password

```txt
User Input
    ↓
cmd/add.go (handleAddManual/handleAddGenerate)
    ↓
Validate service doesn't exist
    ↓
Generate password (if needed)
    ↓
Create PasswordEntry
    ↓
Add to Vault.Entries
    ↓
storage.SaveVault()
    ↓
If vault.Encrypted:
  - Serialize to JSON
  - Encrypt with AES-256-GCM
  - Base64 encode
  - Write to vault.json
Else:
  - Serialize to JSON
  - Write to vault.json
    ↓
User confirmation message
```

### Copying a Password

```text
User Input (service name)
    ↓
cmd/copy.go (handleCopy)
    ↓
storage.LoadVault()
    ↓
If vault encrypted:
  - Read vault.json
  - Base64 decode
  - Decrypt with AES-256-GCM
  - Parse JSON
  - Mark as decrypted
Else:
  - Read vault.json
  - Parse JSON
    ↓
Retrieve password from entries
    ↓
utils.CopyToClipboard()
    ↓
User confirmation message
```

## Design Patterns

### 1. Command Pattern

Each subcommand is a function that implements the command logic. Cobra routes CLI arguments to appropriate handler.

### 2. Separation of Concerns

- **Commands** handle user interaction and validation
- **Storage** handles file I/O and encryption/decryption
- **Crypto** handles cryptographic operations
- **Models** define data structures
- **Utils** provide common utilities

### 3. Error Handling

Go's idiomatic error handling pattern:

```go
func DoSomething() error {
    if err := someOperation(); err != nil {
        return fmt.Errorf("descriptive message: %w", err)
    }
    return nil
}
```

### 4. Immutable Configuration

Encryption parameters are constants defined in `crypto.go`:

```go
const (
    pbkdf2Iterations = 100_000
    saltLen = 16
    nonceLen = 12
    keyLen = 32
)
```

## Security Considerations

### Password Storage

- Passwords are encrypted using AES-256-GCM when vault encryption is enabled
- Passwords are stored only in memory during command execution
- No passwords are logged or printed (except via `open` command intentionally)

### Key Derivation

- Uses PBKDF2-SHA256 with 100,000 iterations
- Random salt prevents rainbow table attacks
- Computationally expensive to brute-force

### Input Security

- Password input is hidden from terminal using `stty` on Unix-like systems
- Windows falls back to unsecured input (limitation)

### File Permissions

- Vault file is created with 0600 permissions (read/write for owner only)

## Dependencies

### Go Standard Library

- `encoding/json` - JSON serialization
- `crypto/aes` - AES cipher
- `crypto/cipher` - GCM mode
- `crypto/rand` - Cryptographically secure randomness
- `crypto/sha256` - SHA-256 hashing
- `os`, `io` - File and stream operations

### External Packages

- `github.com/spf13/cobra` - CLI framework
- `github.com/atotto/clipboard` - Clipboard operations
- `golang.org/x/crypto/pbkdf2` - PBKDF2 key derivation

## Building and Distributing

### Single Binary

Go compiles to a single executable with no runtime dependencies. This makes distribution simple:

- Linux: `go-passman-linux-amd64`
- macOS Intel: `go-passman-darwin-amd64`
- macOS ARM: `go-passman-darwin-arm64`
- Windows: `go-passman-windows-amd64.exe`

### Build Scripts

- `build.sh` - Unix/Linux build script with cross-platform support
- `build.bat` - Windows batch script
- `Makefile` - Standard make commands

## Performance

### Memory Usage

- Vault loaded entirely into memory
- For typical usage (< 10,000 entries), memory footprint is minimal

### Encryption Performance

- PBKDF2 with 100,000 iterations takes ~100ms per encryption/decryption
- Acceptable for interactive CLI usage

### Startup Time

- Go binaries are fast to start
- Typical startup: < 10ms

## Limitations and Future Improvements

### Current Limitations

- Only one user (no multi-user support)
- No synchronization/backup features
- Password input on Windows is not hidden
- Vault must be encrypted/decrypted entirely

### Future Improvements

- [ ] TUI (Text User Interface) using Bubble Tea
- [ ] Database backend (SQLite) instead of JSON
- [ ] Cloud synchronization
- [ ] Password expiration warnings
- [ ] Search/filter capabilities
- [ ] Import/export functionality
- [ ] Web UI option

## Testing

Unit tests should be added in `*_test.go` files alongside implementation files:

- `internal/crypto/crypto_test.go` - Encryption tests
- `internal/storage/storage_test.go` - Storage tests
- `cmd/*_test.go` - Command tests

Example test:

```go
func TestGeneratePassword(t *testing.T) {
    password := utils.GeneratePassword(16, true, true)
    if len(password) != 16 {
        t.Errorf("Expected length 16, got %d", len(password))
    }
}
```

Run tests with:

```bash
go test ./...
go test -v ./...      # Verbose
go test -cover ./...  # Coverage
```
