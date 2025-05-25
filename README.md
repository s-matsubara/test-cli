# test-cli

A sample CLI tool demonstrating modern Go CLI development practices with automated build and distribution.

## Overview

test-cli serves as an example CLI application showcasing:
- Production-grade build automation with GoReleaser
- Multi-platform distribution (Linux, macOS, Windows)
- Professional development tooling with mise
- Comprehensive testing setup
- Automated Homebrew tap integration

## Installation

### Via Homebrew
```bash
brew tap s-matsubara/homebrew-tap
brew install test-cli
```

### From Releases
Download the latest binary from [GitHub Releases](https://github.com/s-matsubara/test-cli/releases).

### From Source
```bash
git clone https://github.com/s-matsubara/test-cli.git
cd test-cli
go build -o test-cli
```

## Usage

```bash
# Basic usage
test-cli

# Show version information
test-cli version

# Help
test-cli --help
```

## Development

### Prerequisites
- Go 1.24.3 or later
- [mise](https://mise.jdx.dev/) for task management (optional)

### Building
```bash
# Using mise
mise run build

# Using go directly
go build -o test-cli
```

### Testing
```bash
# Run all tests
mise run go-test-all
# or
go test ./... -v

# Integration tests
INTEGRATION_TEST=true go test ./...
```

### Linting
```bash
mise run go-lint
# or
golangci-lint run
```

## License

Copyright © 2023 matsubara
