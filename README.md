# Disk Space Analyzer

![CI](https://github.com/Qyroxen/Disk-Space-Analyzer/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Disk-Space-Analyzer/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Disk-Space-Analyzer?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Disk-Space-Analyzer)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Disk-Space-Analyzer)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Disk-Space-Analyzer?style=social)](https://github.com/Qyroxen/Disk-Space-Analyzer/stargazers)

## What is it?

Disk Space Analyzer is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Disk-Space-Analyzer.git
cd Disk-Space-Analyzer
go build -o diskspaceanalyzer .

# Run
./diskspaceanalyzer --help
```

## CLI Usage

```bash
# Basic usage
./diskspaceanalyzer

# With flags
./diskspaceanalyzer --verbose --output json

# Get help
./diskspaceanalyzer --help
```

## Examples

```bash
# Example 1
./diskspaceanalyzer example1

# Example 2
./diskspaceanalyzer example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o diskspaceanalyzer .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Disk-Space-Analyzer/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Disk-Space-Analyzer?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Disk-Space-Analyzer/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Disk-Space-Analyzer?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Disk-Space-Analyzer/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Disk-Space-Analyzer" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Disk-Space-Analyzer/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Disk-Space-Analyzer" alt="Pull Requests">
  </a>
</p>
