# Installation Guide

This guide provides detailed instructions for installing and using go-passman on different platforms.

## Prerequisites

- **Method 1 (Build from Source):** Go 1.19 or higher, Git
- **Method 2 (GitHub Releases):** none (just download and run)
- Basic command-line knowledge

## Installation Methods

### Method 1: Build from Source

#### Windows

1. **Install Go**
   - Download from https://golang.org/doc/install
   - Run the installer and follow the prompts
   - Verify installation:

     ```bash
     go version
     ```

2. **Clone the Repository**

   ```bash
   git clone https://github.com/KravetsMikhail/go-passman.git
   cd go-passman
   ```

3. **Build the Application**

   ```bash
   go build -o go-passman.exe
   ```

   Or use the provided script:

   ```bash
   .\build.bat
   ```

4. **Run go-passman**

   From the project directory use:
   - `.\go-passman.exe --help` — show help
   - `.\go-passman.exe add` — add entry
   - `.\go-passman.exe list` — list entries

   On Windows the executable has the `.exe` extension; use backslash: `.\go-passman.exe` (analogous to `./go-passman` on Linux/macOS).

5. **(Optional) Add to PATH**
   - Move `go-passman.exe` to a directory in your PATH
   - Or create a batch file in a PATH directory that calls the full path
   - Then use `go-passman` from anywhere

#### Linux/macOS

1. **Install Go**
   - Download from https://golang.org/doc/install
   - Extract and add to PATH:

     ```bash
     # Example for Linux
     wget https://golang.org/dl/go1.22.0.linux-amd64.tar.gz
     tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
     export PATH=$PATH:/usr/local/go/bin
     ```

2. **Clone the Repository**
   ```bash
   git clone https://github.com/KravetsMikhail/go-passman.git
   cd go-passman
   ```

3. **Build the Application**
   ```bash
   go build -o go-passman
   ```
   
   Or use the provided script:
   ```bash
   chmod +x build.sh
   ./build.sh linux
   # or for macOS:
   ./build.sh darwin
   ```

4. **Run go-passman**

   From the project directory: `./go-passman --help`, `./go-passman add`, etc. If installed system-wide, use `go-passman` without `./`.

5. **(Optional) Install to System**
   ```bash
   # Copy to system path
   sudo cp go-passman /usr/local/bin/
   
   # Now you can run from anywhere
   go-passman --help
   ```

### Method 2: Download from GitHub Releases

The simplest way to install without building — download a pre-built binary for your platform.

**1. Go to [Releases](https://github.com/KravetsMikhail/go-passman/releases)**

**2. Download the appropriate file:**

| Platform | File |
|----------|------|
| Windows (x64) | `go-passman-windows-amd64.exe` |
| Linux (x64) | `go-passman-linux-amd64` |
| macOS Intel | `go-passman-darwin-amd64` |
| macOS Apple Silicon (M1/M2/M3) | `go-passman-darwin-arm64` |

**3. Install and run**

#### Windows

1. Download `go-passman-windows-amd64.exe`
2. Move to a folder (e.g. `C:\Programs\go-passman\` or your project folder)
3. Rename to `go-passman.exe` (optional)
4. Run from Command Prompt or PowerShell:

   ```batch
   .\go-passman.exe --help
   .\go-passman.exe add
   ```

5. **(Optional)** Add the folder to PATH to run `go-passman` from anywhere

#### Linux

1. Download `go-passman-linux-amd64`
2. Make executable and run:

   ```bash
   chmod +x go-passman-linux-amd64
   ./go-passman-linux-amd64 --help
   ```

3. **(Optional)** Install system-wide:

   ```bash
   sudo mv go-passman-linux-amd64 /usr/local/bin/go-passman
   go-passman --help
   ```

#### macOS

1. Download `go-passman-darwin-amd64` (Intel) or `go-passman-darwin-arm64` (Apple Silicon)
2. Make executable and run:

   ```bash
   chmod +x go-passman-darwin-arm64
   ./go-passman-darwin-arm64 --help
   ```

3. **(Optional)** If macOS blocks the binary:

   ```bash
   xattr -d com.apple.quarantine go-passman-darwin-arm64
   ```

4. **(Optional)** Install system-wide:

   ```bash
   sudo mv go-passman-darwin-arm64 /usr/local/bin/go-passman
   go-passman --help
   ```

**Note:** The vault file (`vault.json`) is created in the **same directory as the executable**. Use `go-passman path` to see where it is stored.

### Method 3: Using make

If you have `make` installed:

```bash
# Build
make build

# Run
make run

# Clean
make clean
```

## Configuration

### Vault File Location

The vault file (`vault.json`) is stored in the **same directory as the go-passman executable**.

**Examples:**
- Windows: `C:\Users\YourName\Programs\go-passman.exe` → vault at `C:\Users\YourName\Programs\vault.json`
- Linux: `/usr/local/bin/go-passman` → vault at `/usr/local/bin/vault.json`
- macOS: `~/Applications/go-passman` → vault at `~/Applications/vault.json`

To find your vault location:

```bash
go-passman path
```

### Backup Your Vault

Since the vault contains sensitive information, keep regular backups:

```bash
# Find your vault
go-passman path

# Back up the vault file
cp vault.json vault.json.backup
```

### Moving the Application

If you move the executable to a different location:

1. The vault file will be created in the new location on first use
2. To move with your vault:

   ```bash
   # Back up vault from old location
   cp /old/location/vault.json ./vault.json
   
   # Move executable and vault to new location
   cp go-passman /new/location/
   cp vault.json /new/location/
   ```

## Upgrading

1. **Back up your vault**

   ```bash
   go-passman path
   cp <vault-path>/vault.json ./vault.json.backup
   ```

2. **Update the binary**

   - **From releases:** Download the new version from [GitHub Releases](https://github.com/KravetsMikhail/go-passman/releases) and replace the old executable
   - **From source:** `git pull origin main` and `go build -o go-passman`

3. **Test new version**

   ```bash
   ./go-passman list
   ```

4. **Keep backup safe**

   ```bash
   cp vault.json.backup <safe-location>/
   ```

## Troubleshooting

### "go: command not found"

- Go is not installed or not in PATH
- Follow the "Install Go" section above
- Verify: `go version`

### "vault.json not found"

- This is normal on first use
- Create first password: `go-passman add`
- Vault will be created automatically

### "Permission denied" (Linux/macOS)

- Make file executable: `chmod +x go-passman`

### "Cannot open vault in editor"

- Ensure editor is installed and in PATH
- Try with absolute path: `go-passman open /usr/bin/nano`

### Password input not hidden on Windows

- This is a known limitation
- Use secure password managers or enter carefully
- Alternative: Use Git Bash or WSL for hidden input

### Vault corrupted

- Check with text editor (if unencrypted)
- Try decrypting if encrypted: `go-passman decrypt`
- Restore from backup if available

## Uninstallation

### Windows

```bash
# Delete the executable
del go-passman.exe

# Delete the vault file (optional)
del vault.json
```

### Linux/macOS

```bash
# Remove from system path
sudo rm /usr/local/bin/go-passman

# Delete vault file (optional)
rm vault.json
```

## Verify Installation

```bash
# Check version
go-passman --version

# List available commands
go-passman --help

# Try adding a test password
go-passman add
# (Enter service name: test)
# (Enter password: test123)

# List passwords
go-passman list

# Remove test entry
go-passman remove
# Select "test" by number from the list, confirm (y)
```

## System Requirements

### Minimum

- 512 MB RAM
- 10 MB disk space for binary
- Additional space for vault file (typically < 1 MB)

### Recommended

- 1 GB RAM
- 100 MB disk space
- Internet connection (for building from source)

## Platform-Specific Notes

### Windows

- Works with Command Prompt and PowerShell
- Password input is visible (not hidden)
- Use `go-passman.exe` or set file association

### Linux

- Works with bash, zsh, and other shells
- Use `sudo` for system-wide installation
- Password input is hidden with `stty`

### macOS

- Intel and Apple Silicon (M1/M2/M3) supported
- Works with Terminal.app or iTerm2
- Password input is hidden
- May need to allow execution:

  ```bash
  xattr -d com.apple.quarantine ./go-passman
  ```

## Advanced: Cross-Platform Building

Build for multiple platforms on one system:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o go-passman-linux

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o go-passman-darwin

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o go-passman-darwin-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o go-passman.exe
```

Or use the provided scripts:

```bash
# bash
./build.sh all

# PowerShell / cmd
build.bat all
```

## Getting Help

- Check the [README.md](README.md) for usage
- Review [ARCHITECTURE.md](ARCHITECTURE.md) for technical details
- See [TESTING.md](TESTING.md) for testing guide
- Open an issue on GitHub for problems

## Security Considerations

1. **Keep vault backed up** - Regular backups are essential
2. **Use strong passwords** - Especially for encryption
3. **Keep go updated** - Security patches are important
4. **Store vault securely** - Keep backups in safe location
5. **Don't share vault** - Keep vault.json private

## Next Steps

After installation:

1. **Create your first password**

   ```bash
   go-passman add --generate
   ```

2. **Encrypt your vault**

   ```bash
   go-passman encrypt
   ```

3. **Verify it works**

   ```bash
   go-passman list
   ```

4. **Back up your vault**

   ```bash
   go-passman path  # Find vault location
   # Back up vault.json to safe location
   ```

Enjoy using go-passman! 🔐
