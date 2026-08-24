# ⚙️ Disk Space Analyzer

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Management tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`management` `operations` `cli` `golang`

---

## What is Disk-Space-Analyzer?

**Disk-Space-Analyzer** is an operations management tool for automating, tracking, and coordinating development workflows.

## Features

- ✅ `humanSize()` — Humansize
- ✅ Formatted output
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Disk-Space-Analyzer.git
cd Disk-Space-Analyzer

# Build
go build -o disk-space-analyzer .

# Run
./disk-space-analyzer [file]
```

### Or directly with `go run`:
```bash
go run main.go [file]
```

## Usage

```bash
# Basic usage
./disk-space-analyzer [file]
```

### Example Output

```
$ ./disk-space-analyzer [file]
Disk Space Analyzer
===================
Top 10 largest files:
```

## Project Structure

```
Disk-Space-Analyzer/
  main.go          # Entry point (56 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
