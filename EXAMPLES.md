# Usage Examples

This document provides practical examples for using go-passman in various scenarios.

## Running the program

- **Windows** (Command Prompt or PowerShell): from the folder with the executable run `.\go-passman.exe` (e.g. `.\go-passman.exe --help`, `.\go-passman.exe add`). If `go-passman.exe` is in your PATH, you can run `go-passman` without path.
- **Linux / macOS**: from the project folder run `./go-passman` (e.g. `./go-passman --help`). If installed in PATH, use `go-passman` alone.

In the examples below, `go-passman` stands for the appropriate invocation on your system.

## Initial Setup

### First Run

```bash
# Show help to see available commands
go-passman --help

# Check where your vault will be stored
go-passman path
# Output: Vault path: /home/user/go-passman/vault.json  (or Windows path)
```

### Add Your First Password

#### Manual Entry

```bash
go-passman add

# Prompts:
# Enter service name: github
# Enter login (optional, press Enter to skip): login1
# Enter host (optional, press Enter to skip): 1.1.1.1
# Enter comment (optional, press Enter to skip):
# Enter password: ••••••••••••••
# (Input is hidden)

# Output: ✅ Password for 'github' saved.
```

#### With Generated Password

```bash
go-passman add --generate

# Prompts:
# Enter service name: aws
# Enter login (optional, press Enter to skip): login1
# Enter host (optional, press Enter to skip): localhost
# Enter comment (optional, press Enter to skip):
# Enter password length (default 16): 24
# Include numbers? (y/n, default y): y
# Include special characters? (y/n, default y): y

# Output: ✅ Password for 'aws' saved and copied to clipboard.
# (Password is automatically copied, ready to paste elsewhere)
```

## Everyday Usage

### View All Passwords

Entries are **numbered** (1, 2, 3…) in the same order every time. You can use the number with `copy N` or `remove N` instead of typing long names.

By default, list is **compact** (one line per entry), so the output fits narrow terminals:

```bash
go-passman list

# Output (compact format):
# 🔐 Saved entries (use copy N or remove N):
#
#   1.  github · john.doe · github.com · work account
#   2.  gmail · jane · mail.google.com · -
#   3.  aws · admin · - · production
#   4.  spotify · - · - · -
```

For a **table** with aligned columns (use when your terminal is wide enough):

```bash
go-passman list -t
# or
go-passman list --table

# Output (table format):
# 🔐 Saved entries (use copy N or remove N):
#
#   #     Service          Login                 Host                     Comment
#   ----  ----------------  --------------------  -----------------------  ----------------------
#   1     github            john.doe              github.com               work account
#   2     gmail             jane                  mail.google.com          -
```

### Copy a Password to Clipboard

You can copy by **service name** or by **number** (same order as in `list`):

```bash
go-passman copy github
# or, if github is 2nd in the list:
go-passman copy 2

# Output: 📋 Password for 'github' copied to clipboard!
# (You can now paste it: Ctrl+V or Cmd+V)
```

### Update an Existing Password

Each field shows the **current value**. Press **Enter** to keep it, or type a new value to replace. After update, the **new values** are printed, then the **resulting list** is shown and you are asked **Continue? (y/n)** — answer **y** to update another entry in the same run (30 second timeout; if no answer, the command exits).

```bash
# Manual update
go-passman update

# Prompts:
# Select a service to update:
# 1. github
# 2. aws
# 3. gmail
# Enter your choice (number): 1
#
# Login [john.doe]: 
# Host [github.com]: 
# Comment [work]: 
# Password (Enter to keep current): ••••••••

# Output:
# ✅ Password for 'github' updated.
#
#   New values:
#     Service: github
#     Login:   john.doe
#     Host:    github.com
#     Comment: work
#     Password: ****
#
# 🔐 Saved entries ...
#   (resulting list)
#
# Continue? (y/n): y   ← answer y to update another, or n to exit (30s timeout)
```

```bash
# Update with generated password (Login/Host/Comment: Enter = keep, type = replace)
go-passman update --generate

# Prompts:
# Select a service to update: ...
# Login [john.doe]: 
# Host [github.com]: 
# Comment [work]: 
# Enter password length (default 16): 20
# Include numbers? (y/n, default y): y
# Include special characters? (y/n, default y): y

# Output: ✅ Password for 'github' updated and copied to clipboard.
#   New values: ...
```

### Remove a Password

Like **update**, you first see a list and select by number. After removal, the **resulting list** is shown and **Continue? (y/n)** lets you remove more in the same run (30 second timeout).

```bash
go-passman remove

# Prompts:
# Select a service to remove:
# 1. github
# 2. gmail
# 3. spotify
# Enter your choice (number): 3
# Are you sure you want to remove 'spotify'? (y/n): y

# Output: ✅ Service 'spotify' removed.
#
# 🔐 Saved entries ...
#   (resulting list)
#
# Continue? (y/n): y   ← remove another, or n to exit (30s timeout)
```

## Security

### Encrypt Your Vault

```bash
go-passman encrypt

# Prompts:
# Enter master password: ••••••••••••••
# Confirm password: ••••••••••••••

# Output: ✅ Vault encrypted successfully.
```

After encryption, every time you use go-passman, it will ask for the master password:

```bash
go-passman list

# Prompts:
# Vault is encrypted. Please enter your password: ••••••••••••••

# Output:
# 🔐 Saved services:
# - github
# - aws
# - gmail
```

### Decrypt Your Vault

If you want to remove encryption (not recommended):

```bash
go-passman decrypt

# Output: ✅ Vault decrypted successfully.
```

### Check Vault Status

```bash
go-passman status

# Output:
# 🔐 Vault Status:
#   Entries: 3
#   Encrypted: true
#   Path: /home/user/go-passman/vault.json
```

## Advanced Usage

### Edit Vault Directly

Open your vault in a text editor:

```bash
# With default editor (cat on Unix)
go-passman open

# With specific editor
go-passman open vim
go-passman open nano
go-passman open notepad  # Windows
```

This opens the vault in the specified editor. You can manually edit entries:

```json
{
  "entries": {
    "github": {
      "login": "login1",
      "host": "localhost:8088",
      "comment": "comment 1",
      "password": "myPassword123!",
      "encrypted": false
    },
    "aws": {
      "login": "login2",
      "host": "localhost:8181",
      "comment": "comment 2",
      "password": "anotherPassword456!",
      "encrypted": false
    }
  },
  "encrypted": false
}
```

Save and close the editor, and go-passman will update the vault.

### Find Your Vault File

```bash
go-passman path

# Output: Vault path: /home/user/go-passman/vault.json
```

You can then:

- Back it up: `cp /home/user/go-passman/vault.json ~/backups/vault.json.backup`
- Share it (with caution): Transfer to another machine
- Inspect it: Open with text editor (if not encrypted)

## Backup and Recovery

### Create a Backup

```bash
# Find vault location
VAULT_PATH=$(go-passman path | cut -d' ' -f3)

# Create backup
cp "$VAULT_PATH" "$VAULT_PATH.backup"

# Verify backup
ls -lh "$VAULT_PATH"*
```

### Restore from Backup

```bash
# If vault is corrupted
VAULT_PATH=$(go-passman path | cut -d' ' -f3)
cp "$VAULT_PATH.backup" "$VAULT_PATH"

# Verify
go-passman status
```

## Workflow Examples

### Developer Workflow

Store all your development credentials:

```bash
# Add GitHub
go-passman add
# Service: github
# Login: your_github_login
# Host: your_github_host
# Comment: your_github_comment
# Password: your_github_token

# Add GitLab
go-passman add
# Service: gitlab
# Login: your_gitlab_login
# Host: your_gitlab_host
# Comment: your_gitlab_comment
# Password: your_gitlab_token

# Add NPM
go-passman add --generate
# Service: npm
# Login: your_npm_login
# Comment: your_npm_comment
# (Use generated password)

# When you need a password
go-passman copy github
# Now paste in terminal or wherever needed
```

### Database Administrator

```bash
# PostgreSQL credentials
go-passman add
# Service: postgres-prod
# Login: postgres_login
# Host: postgres_host
# Comment: postgres_comment
# Password: (secure password)

go-passman add
# Service: postgres-dev
# Login: postgres_login
# Host: postgres_host
# Comment: postgres_comment
# Password: (secure password)

# MySQL credentials
go-passman add --generate
# Service: mysql-backup
# Login: mysql_login
# Host: mysql_host
# Comment: mysql_comment
# (Auto-generate backup password)

# List all database credentials
go-passman list

# Copy password when needed
go-passman copy postgres-prod
```

### Multi-Environment Management

```bash
# Add for each environment
go-passman add --generate
# Service: api-staging-key
# (20 character password with special chars)

go-passman add --generate
# Service: api-prod-key
# (20 character password with special chars)

go-passman add --generate
# Service: api-dev-key
# (16 character password)

# Encrypt everything
go-passman encrypt

# View status
go-passman status
# Encrypted: true, Entries: 3
```

## Password Generation Recipes

### Strong for Online Banking

```bash
go-passman add --generate
# Enter password length: 20
# Include numbers: y
# Include special characters: y
# (Creates strong: ABc!@#x1234...xyz)
```

### Medium for Regular Sites

```bash
go-passman add --generate
# Enter password length: 16
# Include numbers: y
# Include special characters: n
# (Creates: Abc123xyz789def)
```

### Simple for Low-Security Sites

```bash
go-passman add --generate
# Enter password length: 12
# Include numbers: y
# Include special characters: n
# (Creates: AbcXyz123456)
```

## Troubleshooting Examples

### Forgot Master Password?

You cannot recover a master password. You can:

1. Restore vault from backup (if you have one)
2. Create new vault with new master password

```bash
# Back up old vault (encrypted, now useless)
cp vault.json vault.json.encrypted.backup

# New vault will be created automatically on next use
go-passman add
# This creates a new, unencrypted vault

# Encrypt the new vault with new password
go-passman encrypt
```

### Accidentally Deleted Entry?

If you have a backup:

```bash
# Restore from backup
cp vault.json.backup vault.json

# Verify
go-passman list
```

If no backup, the entry is lost.

### Can't Remember Service Name?

```bash
# List all services
go-passman list

# Find the one you need in the list
```

### Vault File Corrupted?

If unencrypted:

```bash
# Edit directly with text editor
go-passman open vim
# Fix JSON syntax errors
# Save and close
```

If encrypted and you forgot password:

```bash
# Restore from backup only option
cp vault.json.backup vault.json
```

## Integration with Other Tools

### Using with Git

Store GitHub token:

```bash
go-passman add --generate
# Service: github-token
```

Copy when needed:

```bash
go-passman copy github-token
# Use token in GitHub URL: https://token@github.com/username/repo
```

### Using with SSH

Store SSH passphrases:

```bash
go-passman add
# Service: ssh-server
# Password: (your passphrase)
```

Copy when prompted:

```bash
ssh user@server.com
# Password prompt appears
# Paste from clipboard: Ctrl+Shift+V (terminal-dependent)
```

### Using with Cloud Services

```bash
# AWS
go-passman add --generate
# Service: aws-access-key

# Azure
go-passman add --generate
# Service: azure-token

# Google Cloud
go-passman add --generate
# Service: gcp-api-key

# Encrypt all
go-passman encrypt
```

## Regular Maintenance

### Weekly Check

```bash
# Verify vault is encrypted
go-passman status

# List all entries
go-passman list

# Ensure backup exists
ls -lh ~/backups/vault.json.backup
```

### Monthly Rotation

Update important passwords monthly:

```bash
# Update GitHub token
go-passman update --generate
# Select: github-token

# Update AWS credentials
go-passman update --generate
# Select: aws-access-key

# Verify updates
go-passman status
```

### Yearly Audit

```bash
# List everything
go-passman list

# Review and remove unused services
go-passman remove
# Select by number, confirm (y); then "Continue? (y/n)" to remove more or n to exit

# Update all passwords
for service in $(go-passman list | grep -v "^📭"); do
  go-passman update --generate
  # Select each service
done

# Create fresh backup
cp vault.json ~/backups/vault.json.yearly-backup
```

## Tips and Tricks

### Alias Creation (bash/zsh)

```bash
# Add to ~/.bashrc or ~/.zshrc
alias pw="go-passman"

# Now you can use shorter commands
pw copy github
pw list
pw add --generate
```

### Script for Frequent Operations

```bash
#!/bin/bash
# quick-password.sh

echo "Select action:"
echo "1. Copy password"
echo "2. Add new"
echo "3. List all"
read -p "Choice: " choice

case $choice in
  1) read -p "Service: " service
     go-passman copy "$service" ;;
  2) go-passman add --generate ;;
  3) go-passman list ;;
esac
```

### Cron Job for Backups

```bash
# Add to crontab -e
0 2 * * * cp /path/to/vault.json ~/backups/vault.json.$(date +\%Y\%m\%d)
```

This backs up your vault daily at 2 AM.
