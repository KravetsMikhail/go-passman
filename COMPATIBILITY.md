# Go Version Compatibility

## Minimum Supported Version

**go-passman** requires **Go 1.19 or higher**.

## Supported Versions

- ✅ Go 1.19
- ✅ Go 1.20
- ✅ Go 1.21
- ✅ Go 1.22
- ✅ Go 1.23 and newer

## What Changed

### Go 1.19 Features Used

- Standard library features (crypto, encoding, etc.)
- Generics (using constraints package)
- Interface type parameters

### Removed Go 1.22+ Features

In order to support Go 1.19, we removed:

- **Range over integer** - Replaced `for i := range length` with traditional `for i := 0; i < length; i++`
- **Range over functions** - Not used in this project
- **Clear builtin** - Not used in this project

## Building for Different Versions

### Verify Your Go Version

```bash
go version
```

### Build with Specific Version

All Go versions from 1.19+ will automatically work:

```bash
go build -o go-passman
```

## Dependencies

All external dependencies are compatible with Go 1.19+:

| Package | Min Go Version | Status |
|---------|----------------|--------|
| github.com/spf13/cobra | 1.0+ | ✅ Compatible |
| github.com/atotto/clipboard | 1.0+ | ✅ Compatible |
| golang.org/x/crypto | 1.0+ | ✅ Compatible |

## Upgrading from Go 1.18 or Earlier

If you have Go 1.18 or earlier, you'll need to upgrade:

1. Download Go 1.19+ from [golang.org](https://golang.org/dl/)
2. Install following the official instructions
3. Verify: `go version` should show 1.19+

## Testing

The project has been tested and verified to work with:

- Go 1.19 (minimum supported)
- Go 1.22+ (latest features available)

## Contributing

When contributing code, please ensure:

- No Go 1.22+ only syntax (like range over integers)
- No newer stdlib features without fallbacks
- Test with Go 1.19 where possible

## Troubleshooting

### Error: "version 1.22 required"

This means your go.mod file specifies Go 1.22. Update it:

```bash
go mod edit -go=1.19
```

### Compilation Errors with Older Go

If you have Go < 1.19, upgrade Go to version 1.19 or newer:

```bash
go version  # Check your version
```

---

**Last Updated**: 2024-01-24  
**Minimum Go Version**: 1.19  
**Tested with**: Go 1.19, 1.20, 1.21, 1.22+
