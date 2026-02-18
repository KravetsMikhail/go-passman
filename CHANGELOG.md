# Changelog

All notable changes to go-passman will be documented in this file.

## [Unreleased]

## [0.3.1] - 2026-02-17

### Added

- **Web UI inactivity timer**: if the user is inactive for N minutes (no mouse, keyboard, scroll), the session is locked and the user is redirected to the unlock page. Default: 5 minutes. Set `INACTIVITY_MINUTES` env var to change (e.g. `INACTIVITY_MINUTES=10 go-passman -w`).

## [0.3.0] - 2026-02-13

### Added

- **-w / --web**: run go-passman as a web server with a simple web UI. List, add, edit, delete entries and view password. Encrypted vault: unlock once in the browser. Server listens on 127.0.0.1:8080 (set `WEB_PORT` to change port).
- **list -f / --filter**: filter output by name (substring, case-insensitive). Example: `list -f git` shows only entries whose name contains "git". Displayed numbers match the full list so `copy N` works as expected.
- **list**: "Найдено X из Y" in the header (found X of Y entries).
- **Update** and **Remove**: after each action the resulting list is shown; question **Continue? (y/n)** allows multiple updates or removals in one run. **30 second timeout** on "Continue?" — if no answer, the command exits (vault no longer held in memory).
- **Add** (encrypted vault): vault password is asked **first**, then service name and other fields.
- **Large vaults (100+ entries)**:
  - **list**: paginated by 20 entries per page; "Press Enter for next page (q to quit)". When stdout is not a terminal (e.g. `list > file`), all entries are printed without pagination.
  - **update** / **remove**: optional **filter by name** (substring, Enter = show all); when more than 25 entries, choice is paginated (n = next page, q = quit). Cancelling with q exits the update/remove loop.
- **Documentation**: explicit description of **running the program on Windows** (`.\go-passman.exe` from the project folder, or `go-passman` if in PATH); Linux/macOS `./go-passman`. README, QUICKSTART, INSTALL, EXAMPLES, and 00_START_HERE updated with Windows invocation and recent behavior.

## [0.2.0] - 2026-02-07

### Major Changes

- Added optional fields: **login**, **host**, **comment** for each entry
- **List command**: compact format by default (one line per entry, fits narrow terminals); optional table view with `list -t` / `list --table`
- List shows Service · Login · Host · Comment; empty fields shown as `-`
- **Numbered entries** in list (1, 2, 3…); **copy N** — copy by number (e.g. `copy 2`). **Remove** — like update: list is shown, select entry by number to remove
- **Update command**: no more y/n prompts; each field shows **current value** — press **Enter** to keep, type to replace; **new values** printed after update
- Hidden password input (cross-platform via `golang.org/x/term`)

### Fixed Issues

- ✅ Password input is hidden when entering (no echo; works on Windows, Linux, macOS)
- ✅ Terminal echo restored when pressing Ctrl+C during password prompt (no stuck invisible input)
- ✅ Encrypted vault: password is passed correctly for add, update, remove (no "password required" error after entering once)

## [0.1.0] - 2026-01-24

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

### 0.3.1 (2026-02-17)

### 0.3.0 (2026-02-13)

- Web UI: `-w` / `--web` runs server at http://127.0.0.1:8080
- list: filter `-f`, "Найдено X из Y", pagination for large vaults
- update/remove: multiple actions per run, filter and pagination, 30s timeout on Continue?
- add: vault password first when encrypted
- Docs: Windows invocation (.\go-passman.exe)

### 0.2.0 (2026-02-07)

- Added fields: host and comment (optional)
- Expanded data display in copy
- Fixed password display when entering (stealthy input)
- Fixed errors when entering into update and add

### 0.1.0 (2026-01-24)

- Initial release in Go
- Complete password management system
- Excellent portability and maintainability
