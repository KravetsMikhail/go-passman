# Quick Start Guide

Get up and running with go-passman in 5 minutes.

## Installation (Pick One)

### Option 1: Build from Source (Recommended)

**Windows:**

```bash
# Clone the repository
git clone https://github.com/KravetsMikhail/go-passman.git
cd go-passman

# Build
go build

# Run
.\go-passman.exe --help
```

**Linux/macOS:**

```bash
# Clone the repository
git clone https://github.com/KravetsMikhail/go-passman.git
cd go-passman

# Build
go build -o go-passman

# Run
./go-passman --help
```

### Option 2: Use Build Script

**Windows:**

```bash
build.bat
.\dist\go-passman-windows-amd64.exe --help
```

**Linux/macOS:**

```bash
chmod +x build.sh
./build.sh linux
./dist/go-passman-linux-amd64 --help
```

## First Time Setup

### 1. Create Your First Password

```bash
go-passman add --generate
# Enter service name: github
# Enter login (optional, press Enter to skip): login1
# Enter password length: 20
# Include numbers: y
# Include special characters: y
```

### 2. List Your Passwords

```bash
go-passman list
```

Output:

```bash
🔐 Saved services:
- github
```

### 3. Copy a Password

```bash
go-passman copy github
# Password copied to clipboard!
```

### 4. Encrypt Your Vault (Recommended)

```bash
go-passman encrypt
# Enter master password: ••••••••
# Confirm password: ••••••••
```

### 5. Check Everything

```bash
go-passman status
```

Output:

```bash
🔐 Vault Status:
  Entries: 1
  Encrypted: true
  Path: /path/to/vault.json
```

## Common Tasks

### Add Another Password

```bash
go-passman add
# Enter service name: github
# Enter login (optional, press Enter to skip): login1
# Enter password: ••••••••
```

### Update a Password

```bash
go-passman update --generate
# Select service: github
# (New password is generated and copied)
```

### Remove a Password

```bash
go-passman remove github
# Are you sure? (y/n): y
```

### Back Up Your Vault

```bash
# Find your vault
go-passman path
# Output: Vault path: /home/user/go-passman/vault.json

# Back it up
cp /home/user/go-passman/vault.json ./vault-backup.json
```

## Tips

1. **Encrypted vault is safer** - Always encrypt after setup
2. **Generate strong passwords** - Use the `--generate` flag
3. **Copy instead of typing** - Use `go-passman copy <service>`
4. **Regular backups** - Back up your vault.json file
5. **Remember master password** - You can't recover it!

## Getting Help

| Task | Command |
|------|---------|
| See all commands | `go-passman --help` |
| Help for a command | `go-passman add --help` |
| Show vault location | `go-passman path` |
| Show vault status | `go-passman status` |

## Troubleshooting

### "vault.json not found"

- This is normal on first run
- Will be created when you add first password
- No action needed

### "Cannot find go-passman command"

- Make sure you built it: `go build -o go-passman`
- Or use full path: `./go-passman --help`

### "Vault is encrypted"

- Enter your master password when prompted
- Use the password you set with `go-passman encrypt`

## Next Steps

1. **Read the full guide**: See [README.md](README.md)
2. **Learn advanced features**: Check [EXAMPLES.md](EXAMPLES.md)
3. **Understand the architecture**: Read [ARCHITECTURE.md](ARCHITECTURE.md)
4. **Get help installing**: See [INSTALL.md](INSTALL.md)

## Security Reminder

```bash
# ✅ DO THIS
✓ Encrypt your vault
✓ Use strong master password
✓ Back up your vault.json
✓ Keep backups in safe place
✓ Change passwords regularly

# ❌ DON'T DO THIS
✗ Share your vault.json
✗ Use weak master password
✗ Forget to back up
✗ Store backup unsecurely
✗ Store master password in plain text
```

## Quick Reference

```bash
# Core Commands
go-passman add              # Add password manually
go-passman add -g          # Add with generated password
go-passman copy SERVICE    # Copy to clipboard
go-passman list            # Show all services
go-passman update          # Update password
go-passman remove SERVICE  # Remove password

# Encryption
go-passman encrypt         # Encrypt vault
go-passman decrypt         # Decrypt vault

# Info
go-passman status          # Show vault info
go-passman path            # Show vault location
go-passman --help          # Show all commands

# Editing
go-passman open            # Open in default editor
go-passman open vim        # Open in specific editor
```

## One-Minute Setup

```bash
# 1. Build (30 seconds)
go build -o go-passman

# 2. Add first password (15 seconds)
go-passman add --generate

# 3. Encrypt (15 seconds)
go-passman encrypt

# Done! 💪
go-passman status
```

---

**Ready?** Start with `go-passman add --generate` and you're all set!

For questions, check the [README.md](README.md) or [EXAMPLES.md](EXAMPLES.md).
