# Releaser Library - Copilot Instructions

## Overview

This is a Go library that handles string formatting and transformation for Defacto2 release group names. It provides functions to clean, format, and transform releaser names for various contexts (database cells, URLs, human-readable titles, HTML links).

## Build, Test, and Lint Commands

### Development Tools Required
- **Go** 1.25.2 or later
- **Task** (task runner) - https://taskfile.dev/installation/
- **golangci-lint** (for linting)
- **gofumpt** (for code formatting)

### Common Commands
```bash
# View all available tasks
task --list

# Run all tests
task test

# Run tests with race detection (slower but catches race conditions)
task testr

# Format code and lint
task lint

# Run static analysis for nil dereferences
task nil

# Update dependencies
task update     # Update all dependencies
task patch      # Update only patch versions

# Generate and browse documentation locally
task doc        # Opens pkgsite at localhost:8090
```

### Running Individual Tests
```bash
# Run a single test function
go test -v -run TestCell ./...

# Run tests in a specific package
go test -v ./fix
go test -v ./name
go test -v ./initialism

# Run example tests only
go test -run Example ./...
```

## Architecture

### Package Structure
The library is organized into 4 packages:

#### `releaser` (main package)
- **Public API** - Contains 6 transformation functions
- `Clean(s string)` - Stylizes and cleans strings for display
- `Cell(s string)` - Formats for database table cells (uppercase)
- `Humanize(path string)` - Converts URL paths to human-readable names
- `Link(path string)` - Formats paths as link descriptions (uses `+` instead of `,`)
- `Obfuscate(s string)` - Converts clean names to URL-safe paths
- `Title(s string)` - Formats for titles with acronym deobfuscation
- `Index(path string)` - Converts paths to database index format (uppercase)

#### `name` package
- **URL path handling** - Manages the `Path` type representing URL paths
- Validates path format and retrieves well-known styled names
- Maps between paths and their canonical names
- `Humanize()` - Expands URL paths to full names
- `Obfuscate()` - Converts names to URL-safe paths (slug format)
- Contains curated maps of special names and known releasers

#### `fix` package
- **String manipulation utilities** - Low-level character and string operations
- `StripChars()` - Removes incompatible characters (keeps: A-Z a-z À-Ö Ø-ö ø-ÿ 0-9 - , &)
- `TrimThe()` - Removes leading "The " prefix (for BBS/FTP sites)
- `TrimSP()` - Trims single special characters (/, *)
- `Cell()` - Converts to uppercase for database cells
- `Format()` - Applies title case to the string
- `Abbreviation()` - Handles special acronyms and ordinal numbers

#### `initialism` package
- **Alternative names database** - Maps URLs to acronyms, initialisms, and alternative spellings
- Example: `"acid-productions"` → `["ACiD", "ACiD Prods", "ACiD Productions"]`
- Used by main functions to recognize and transform abbreviated names

### String Transformation Flow

**Clean/Display paths:**
1. Strip incompatible characters
2. Remove leading "The "
3. Trim whitespace
4. Apply title case

**Obfuscate to URL paths:**
1. Check special names map
2. Check initialisms/acronyms map
3. Fall back to manual slug conversion

**Humanize from URL paths:**
1. Look up in special names (returns direct replacement)
2. Otherwise expand using initialisms/names maps
3. Apply Clean() formatting

## Key Conventions

### Character Compatibility
- **Allowed:** A-Z, a-z, À-Ö, Ø-ö, ø-ÿ (accented letters), 0-9, hyphen `-`, comma `,`, ampersand `&`
- **Replaced during cleanup:** Most punctuation and special characters are stripped
- **URL path format:** Lowercase with hyphens (`-`) separating words, underscores (`_`) for special characters like apostrophes

### Naming Patterns

#### Multiple Groups
- **Comma** `,` - Separate distinct groups (e.g., `"TDT, TRSi"`)
- **Asterisk** `*` - Cooperation/collaboration (e.g., `"Class * Paradigm * Razor 1911"`)
- **Slash** `/` - Alternative representation (e.g., `"TDT / TRSi"`)

#### Special Handling
- **"The" prefix** - Automatically stripped from BBS/FTP site names
- **Acronyms/Initialisms** - Preserved and expanded (e.g., `"NAPPA"` → `"North American Pirate-Phreak Association"`)
- **Ampersand handling** - Can be literal `&` or written as `ampersand` in URLs

### Testing Conventions
- Tests use table-driven format with subtests (`t.Run`)
- Example functions in tests demonstrate public API usage
- All public functions have corresponding test functions and examples
- The test file imports the parent module as `releaser_test` to test the public API

### Linting Configuration
- Uses golangci-lint with Go 1.25.2 compatibility
- Disabled checks: `depguard`, `funlen`, `nlreturn`, `varnamelen`, `wsl`, `wsl_v5`
- Code formatters: gci, gofmt, gofumpt, goimports
- Exclusions for generated code, third_party, builtin, and examples directories
- Misspell linting uses US English locale

### Documentation Style
- Functions have detailed comment blocks with examples
- Example comments show input → output format
- Comments explain the transformation pipeline (what happens in order)
