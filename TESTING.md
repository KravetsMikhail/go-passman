# Testing Guide

This guide provides instructions for testing go-passman functionality.

## Manual Testing

### 1. Help Commands

```bash
# Show main help
go-passman --help

# Show version
go-passman --version

# Get help for specific command
go-passman add --help
go-passman copy --help
```

### 2. Basic Operations

#### Test: Add Password Manually

```bash
go-passman add
# Enter service name: github
# Enter login (optional, press Enter to skip): login1
# Enter password: mySecurePassword123
# Expected: ✅ Password for 'github' saved.
```

#### Test: List Passwords

```bash
go-passman list
# Expected: Shows "github" in the list
```

#### Test: Copy Password

```bash
go-passman copy github
# Expected: 📋 Password for 'github' copied to clipboard!
```

#### Test: Show Vault Path

```bash
go-passman path
# Expected: Displays path like: D:\PROJECTS\my\go-passman\vault.json
```

#### Test: Show Vault Status

```bash
go-passman status
# Expected: Shows entries count, encryption status, and path
```

### 3. Password Generation

#### Test: Add with Generated Password

```bash
go-passman add -g
# or
go-passman add --generate
# Enter service name: github-generated
# Enter login (optional, press Enter to skip): login1
# Enter password length (default 16): 20
# Include numbers? (y/n, default y): y
# Include special characters? (y/n, default y): y
# Expected: ✅ Password for 'github-generated' saved and copied to clipboard.
```

### 4. Update Operations

#### Test: Update Password (Manual)

```bash
go-passman update
# Select service from list
# Enter new password: newPassword456
# Expected: ✅ Password for 'service' updated.
```

#### Test: Update with Generation

```bash
go-passman update --generate
# or
go-passman update -g
# Select service, choose options
# Expected: ✅ Password for 'service' updated and copied to clipboard.
```

### 5. Remove Operations

#### Test: Remove Entry

```bash
go-passman remove github
# Are you sure you want to remove 'github'? (y/n): y
# Expected: ✅ Service 'github' removed.
```

#### Test: Remove Non-existent Entry

```bash
go-passman remove nonexistent
# Expected: ❌ Service 'nonexistent' not found.
```

### 6. Encryption/Decryption

#### Test: Encrypt Vault

```bash
go-passman encrypt
# Enter master password: myMasterPassword
# Confirm password: myMasterPassword
# Expected: ✅ Vault encrypted successfully.
```

#### Test: Check Encrypted Status

```bash
go-passman status
# Expected: Encrypted: true
```

#### Test: Load Encrypted Vault

```bash
go-passman list
# Vault is encrypted. Please enter your password: myMasterPassword
# Expected: Shows services
```

#### Test: Decrypt Vault

```bash
go-passman decrypt
# Expected: ✅ Vault decrypted successfully.
```

#### Test: Encrypt with Mismatched Passwords

```bash
go-passman encrypt
# Enter master password: password1
# Confirm password: password2
# Expected: passwords do not match (error)
```

### 7. Editor Integration

#### Test: Open in Default Editor

```bash
go-passman open
# (opens vault in default editor or cat)
# Make changes if needed
# Save and close
# Expected: ✅ Vault updated.
```

#### Test: Open with Specific Editor

```bash
go-passman open notepad
# (opens vault in notepad on Windows)
# Expected: Vault opens in specified editor
```

### 8. Error Cases

#### Test: Incorrect Password for Encrypted Vault

1. Encrypt vault with password "correct"
2. Run any command and enter wrong password
3. Expected: Error message about incorrect password

#### Test: Corrupted Vault File

1. Manually edit vault.json with invalid JSON
2. Run `go-passman list`
3. Expected: Error message

#### Test: Missing Service

```bash
go-passman copy nonexistent
# Expected: ❌ Service 'nonexistent' not found.
```

## Automated Testing

### Run Go Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/crypto/...
```

### Example Test Cases

#### Encryption/Decryption Tests

- Encrypt and decrypt the same data
- Test with special characters
- Test with empty passwords
- Test with very long passwords

#### Password Generation Tests

- Verify correct length
- Test with/without numbers
- Test with/without special characters
- Verify randomness

#### Storage Tests

- Load non-existent vault
- Save and load vault
- Handle encrypted/unencrypted vaults
- File permission checks

## Performance Testing

### Measure Startup Time

```bash
# Linux/macOS
time go-passman list

# Windows
Measure-Command { .\go-passman.exe list }
```

### Measure Encryption Time

```bash
# Add many entries
for i in {1..100}; do
  go-passman add
  # Enter service and password
done

# Time encryption
time go-passman encrypt
```

## Cross-Platform Testing

Test on:

- Windows (cmd, PowerShell)
- Linux (bash, zsh)
- macOS (bash, zsh)

Verify:

- Binary runs without errors
- Clipboard operations work
- Password input is hidden (Unix-like systems)
- Vault file location is correct

## Browser-Based Testing (if web UI added)

- Test all CRUD operations through web interface
- Verify encryption works
- Test across different browsers
- Check responsive design

## Regression Testing

Before each release:

1. [ ] Create fresh vault
2. [ ] Add 10+ passwords
3. [ ] Encrypt vault
4. [ ] Decrypt vault
5. [ ] Update several entries
6. [ ] Remove entries
7. [ ] Copy entries to clipboard
8. [ ] Open in editor and modify
9. [ ] Verify vault integrity

## Known Issues and Limitations

1. **Windows Password Input**: Password input is not hidden on Windows
   - Workaround: Use Ctrl+C to cancel if needed
   - Planned: Better Windows support

2. **Single Encryption Key**: Entire vault uses single master password
   - Different entries cannot have different passwords
   - By design (simplified UX)

3. **No Synchronization**: Vault is stored locally only
   - Planned: Cloud sync feature

4. **No Password Strength Meter**: Generated passwords are always same rules
   - Users control length and character types

## Testing Checklist

- [ ] Help displays correctly
- [ ] Add password (manual and generated)
- [ ] List shows correct entries
- [ ] Copy to clipboard works
- [ ] Update entries
- [ ] Remove entries with confirmation
- [ ] Encrypt vault
- [ ] Decrypt vault
- [ ] Open in editor
- [ ] Show vault path
- [ ] Show vault status
- [ ] Error handling for invalid input
- [ ] Error handling for missing entries
- [ ] Cross-platform compatibility
- [ ] No leftover temp files
- [ ] Vault permissions are correct (0600)

## Continuous Integration

For CI/CD pipelines, use:

```bash
# Format check
go fmt ./...

# Lint
golangci-lint run ./...

# Tests
go test -v ./...

# Coverage
go test -cover ./...

# Build
go build -o go-passman
```

## Help and Support

For issues found during testing:

1. Note exact steps to reproduce
2. Check error messages
3. Review vault.json for integrity
4. Check file permissions
5. Open issue with reproduction steps
