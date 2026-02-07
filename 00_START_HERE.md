# 🎉 Welcome to go-passman!

**go-passman** is a secure, lightweight command-line password manager written in Go.

## ⚡ Quick Start (Choose Your Path)

### 🚀 I Just Want to Use It

1. Read [QUICKSTART.md](QUICKSTART.md) (5 minutes)
2. Build: `go build -o go-passman` (Linux/macOS) or `go build -o go-passman.exe` (Windows)
3. Run: `./go-passman add --generate` (Linux/macOS) or `.\go-passman.exe add --generate` (Windows)

### 📚 I Want to Learn Everything

1. Read [README.md](README.md) for complete documentation
2. Check [EXAMPLES.md](EXAMPLES.md) for practical usage
3. Use [DOCS_INDEX.md](DOCS_INDEX.md) for navigation

### 👨‍💻 I Want to Contribute

1. Read [ARCHITECTURE.md](ARCHITECTURE.md) for code structure
2. Review [TESTING.md](TESTING.md) for test procedures

---

## 📋 What This Project Includes

```
✅ Password Management
   - Add, remove, update, and list passwords
   - Generate strong random passwords
   - Copy to clipboard with one command

✅ Security
   - AES-256-GCM encryption
   - PBKDF2-SHA256 key derivation
   - Secure password input (hidden on Unix-like systems)

✅ Portability
   - Single executable, no dependencies
   - Works on Windows, Linux, and macOS
   - Vault stored in same directory as executable

✅ Documentation
   - User guides (README, QUICKSTART, INSTALL)
   - Technical documentation (ARCHITECTURE)
   - Usage examples (EXAMPLES)
   - Testing guide (TESTING)
```

---

## 🏃 Get Running in 60 Seconds

**Linux / macOS:** `./go-passman` — run from the project folder after `go build -o go-passman`.

**Windows:** `.\go-passman.exe` — run from the project folder after `go build -o go-passman.exe`. If the executable is in PATH, you can use `go-passman` alone.

```bash
# 1. Build (15 seconds)
go build -o go-passman          # or: go build -o go-passman.exe (Windows)

# 2. Add your first password (20 seconds)
./go-passman add --generate     # Windows: .\go-passman.exe add --generate

# 3. Encrypt your vault (15 seconds)
./go-passman encrypt           # Windows: .\go-passman.exe encrypt

# 4. Done! Check status (10 seconds)
./go-passman status            # Windows: .\go-passman.exe status
```

---

## 📖 Documentation Map

| Need | Read |
|------|------|
| Quick setup | [QUICKSTART.md](QUICKSTART.md) |
| All features | [README.md](README.md) |
| Installation help | [INSTALL.md](INSTALL.md) |
| How to use | [EXAMPLES.md](EXAMPLES.md) |
| How it works | [ARCHITECTURE.md](ARCHITECTURE.md) |
| All docs listed | [DOCS_INDEX.md](DOCS_INDEX.md) |
| Need testing info | [TESTING.md](TESTING.md) |
| Project overview | [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md) |

---

## 🎯 Core Commands

```bash
# Add passwords
go-passman add                 # Manual entry
go-passman add --generate      # Auto-generated

# Manage passwords
go-passman list                # List all (numbered; copy N; remove = select from list)
go-passman copy github         # Copy by name or number (e.g. copy 2)
go-passman update              # Update (current value shown; Enter=keep)
go-passman remove              # Select from list, then delete

# Vault security
go-passman encrypt             # Lock vault
go-passman decrypt             # Unlock vault
go-passman status              # Show info
go-passman path                # Show location
```

---

## 🔐 Security First

```bash
# Add passwords
go-passman add

# Encrypt your vault ASAP
go-passman encrypt

# Check status
go-passman status
# Output: Encrypted: true ✅
```

**Remember**: 

- ✅ Encrypt your vault
- ✅ Use strong master password
- ✅ Back up your vault.json
- ✅ Keep backup in safe place

---

## 🏗️ Project Structure

```bash
go-passman/
├── 📄 Documentation
│   ├── 00_START_HERE.md          ← You are here
│   ├── QUICKSTART.md              ← Fast setup
│   ├── README.md                  ← Main guide
│   ├── INSTALL.md                 ← Installation
│   ├── EXAMPLES.md                ← Usage examples
│   └── More docs...
│
├── 💻 Source Code
│   ├── main.go                    ← Entry point
│   ├── cmd/                       ← Commands
│   │   ├── add.go
│   │   ├── remove.go
│   │   ├── copy.go
│   │   ├── list.go
│   │   └── ... (10 commands)
│   │
│   └── internal/                  ← Internal packages
│       ├── crypto/                ← Encryption
│       ├── storage/               ← File I/O
│       ├── models/                ← Data structures
│       └── utils/                 ← Utilities
│
├── 🔧 Build
│   ├── go.mod / go.sum            ← Dependencies
│   ├── build.sh                   ← Unix build
│   ├── build.bat                  ← Windows build
│   └── Makefile                   ← Build targets
│
└── 📋 Info
    ├── LICENSE                    ← MIT
    ├── CHANGELOG.md               ← Version history
    └── ... (more docs)
```

---

## 🚀 Features at a Glance

| Feature | Status | Details |
|---------|--------|---------|
| Add passwords | ✅ | Manual or generated |
| Remove passwords | ✅ | With confirmation |
| Copy to clipboard | ✅ | One command copy |
| List all | ✅ | Show all services |
| Update entries | ✅ | Change existing |
| Encrypt vault | ✅ | AES-256-GCM |
| Decrypt vault | ✅ | Remove encryption |
| Open in editor | ✅ | Edit manually |
| Generate passwords | ✅ | Customizable length |
| Show status | ✅ | Vault information |

---

## 💡 Key Features Explained

### 🔐 Encryption

Your vault is encrypted using AES-256-GCM (military-grade encryption)

```bash
go-passman encrypt
# Now all passwords are encrypted with master password
```

### 🎲 Password Generation

Create strong passwords automatically

```bash
go-passman add --generate
# Enter options, password is generated and copied
```

### 📋 Clipboard Integration

Copy passwords to paste elsewhere

```bash
go-passman copy github
# Password copied! Ready to paste.
```

### 📦 Portable

Vault file is in same directory as executable

```bash
go-passman path
# Vault at: ./vault.json (right here!)
```

---

## 🎓 Learning Resources

### Beginner

1. [QUICKSTART.md](QUICKSTART.md) - Get running fast
2. [README.md](README.md) - Learn the basics
3. Try commands: `go-passman add`, `go-passman list`

### Intermediate

1. [EXAMPLES.md](EXAMPLES.md) - Practical workflows
2. [README.md](README.md) - Deep dive features
3. Try: password generation, encryption, backups

### Advanced

1. [ARCHITECTURE.md](ARCHITECTURE.md) - How it works
2. [TESTING.md](TESTING.md) - Testing procedures
3. Contribute code improvements

---

## ❓ Common Questions

**Q: Where is my vault file stored?**  
A: Run `go-passman path` to find it. It's in the same directory as the executable.


**Q: What encryption does it use?**  
A: AES-256-GCM with PBKDF2-SHA256 key derivation (industry standard).

**Q: How do I back up my passwords?**  
A: Copy the vault.json file to a safe location. See [EXAMPLES.md](EXAMPLES.md).

**Q: What if I forget my master password?**  
A: You cannot recover it. Use your backup if available.

**Q: Is it safe?**  
A: Yes! It uses military-grade encryption and best practices. Always encrypt!

---

## 🛠️ System Requirements

- **Go 1.19 or higher** (if building from source)
- **Windows, Linux, or macOS**
- **~10 MB disk space** for binary
- **~1 MB** per 1,000 passwords

---

## 📞 Next Steps

1. **Choose your path above** ☝️
2. **Read the appropriate guide**
3. **Start using go-passman!**

---

## 📝 Quick Reference Card

```bash
# Setup
go build -o go-passman           # Build
go-passman --help                # Show help
go-passman path                  # Find vault

# Daily use
go-passman list                  # List all (compact); list -t for table
go-passman copy SERVICE          # Copy password
go-passman add --generate        # Add new

# Security
go-passman encrypt               # Encrypt vault
go-passman decrypt               # Remove encryption
go-passman status                # Check status

# Management
go-passman update                # Update password
go-passman remove                 # Select from list; then "Continue?" for more (30s timeout)
go-passman open vim              # Edit manually
```

---

## 🎯 Your First 5 Minutes

```bash
# 1. Build (30 sec)
go build -o go-passman

# 2. Add password (1 min)
go-passman add --generate
# Service: github
# Login: login1
# Host: localhost
# Comment: comment
# Length: 20
# Numbers: y
# Special: y

# 3. List (10 sec)
go-passman list

# 4. Copy password (10 sec)
go-passman copy github
# (Now paste it somewhere!)

# 5. Encrypt (1 min)
go-passman encrypt
# Password: (set master password)
# Confirm: (repeat)

# Done! 🎉
```

---

## ✨ You're All Set!

You now have:

- ✅ A secure password manager
- ✅ Encrypted vault with master password
- ✅ One password safely stored and copied
- ✅ Full documentation available

**Next**: Check [EXAMPLES.md](EXAMPLES.md) for more workflows or [README.md](README.md) for complete documentation.

---

**Start here**: [QUICKSTART.md](QUICKSTART.md)  
**All docs**: [DOCS_INDEX.md](DOCS_INDEX.md)  
**Need help**: Check the relevant guide above

Enjoy using go-passman! 🔐
