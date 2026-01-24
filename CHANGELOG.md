# Changelog

All notable changes to go-passman will be documented in this file.

## [0.1.0] - 2024-01-24

### Major Changes

- **Implementation**: Built in Go
  - Simplified codebase
  - Fast compilation times
  - Straightforward syntax and error handling
  - Excellent cross-platform binary distribution

### Features

- ✨ Command-line interface using Cobra framework
- ✨ Complete password management system
- ✨ Vault file stored in same directory as executable
  - Makes the application more portable
  - Easier to locate and backup vault file
  - Suitable for USB stick or portable applications

### Architecture

- Idiomatic Go `cmd/` and `internal/` packages structure
- Minimal dependency management:
  - 2 external packages + Go standard library
- AES-256-GCM encryption with PBKDF2-SHA256 key derivation (100,000 iterations)

### Build System

- Uses `go.mod` / `go.sum` for dependency management
- Single binary output (no runtime dependencies)
- `build.sh` for Unix-like systems
- `build.bat` for Windows
- Makefile targets for common tasks

### Documentation

- Comprehensive README.md with usage instructions
- ARCHITECTURE.md with technical documentation
- INSTALL.md with platform-specific installation instructions
- TESTING.md with testing procedures
- EXAMPLES.md with practical usage examples

### Fixed Issues

- ✅ Vault location is now explicit and configurable
- ✅ Cross-platform support improved
- ✅ Error messages are more descriptive

### Known Limitations

- Password input is not hidden on Windows (known limitation of Go standard library)
- Workaround: Use WSL or Git Bash on Windows for hidden input

## Compatibility

### Vault Format

- ✅ Secure JSON structure for vault.json
- ✅ AES-256-GCM encryption algorithm
- ✅ PBKDF2-SHA256 key derivation

### Commands

All commands are available:

- `add` - Add new password
- `remove` - Remove password
- `copy` - Copy password to clipboard
- `list` - List all services
- `update` - Update existing password
- `open` - Open vault in editor
- `encrypt` - Encrypt vault
- `decrypt` - Decrypt vault
- `status` - Show vault status
- `path` - Show vault location

## Dependencies

### Go Standard Library

- `encoding/json` - JSON handling
- `crypto/aes` - AES cipher
- `crypto/cipher` - GCM mode
- `crypto/rand` - Cryptographic randomness
- `crypto/sha256` - SHA-256 hashing
- `os`, `io` - File and I/O operations
- `exec` - External command execution

### External Packages

- `github.com/spf13/cobra` - CLI framework
- `github.com/atotto/clipboard` - Clipboard access
- `golang.org/x/crypto/pbkdf2` - PBKDF2 key derivation

## Upgrade Instructions

### Update to Latest Version

```bash
# 1. Backup vault
go-passman path  # Find location
cp <vault-location>/vault.json ./vault.json.backup

# 2. Pull latest changes
git pull origin main

# 3. Rebuild
go build -o go-passman

# 4. Test
./go-passman status
```

## Performance

- Fast startup time (Go binaries are optimized)
- Small binary size (single executable, no runtime)
- Quick compilation

## Future Plans

### Planned Features

- [ ] TUI (Text User Interface) using Bubble Tea
- [ ] SQLite backend for larger vaults
- [ ] Password strength meter
- [ ] Search/filter capability
- [ ] Import/export functionality
- [ ] Cloud sync support
- [ ] Web UI option
- [ ] Multi-user support
- [ ] Password expiration warnings
- [ ] Batch operations

### Technical Debt

- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Improve error handling
- [ ] Add logging capability
- [ ] Performance optimization for large vaults

## Security Considerations

- AES-256-GCM encryption algorithm
- PBKDF2-SHA256 key derivation (100,000 iterations)
- Vault files have 0600 permissions (owner read/write only)
- Master password is never stored or logged

## Testing

- ✅ All manual testing scenarios completed
- ✅ Encryption/decryption verified working
- ✅ Password generation tested
- ✅ Clipboard operations verified
- ✅ Cross-platform compatibility confirmed

## License

MIT License - See LICENSE file for details

---

## Version History

### 0.1.0 (2024-01-24)

- Initial release in Go
- Complete password management system
- Excellent portability and maintainability
