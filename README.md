# 3d-printers

Go binaries for 3D printer control, with self-update from GitHub Releases.

**Docs & Downloads:** https://joeblew999.github.io/3d-printers/

## Binaries

| Binary | Description |
|--------|-------------|
| `x1ctl` | CLI for Bambu Lab X1 printers |
| `fakeprinter` | Demo/test printer server |

## Quick Start

```sh
# Download from releases or install with task
task install:remote NAME=x1ctl

# Or build locally
task build:local
task install:local NAME=x1ctl

# Check version and self-update
x1ctl version --check
x1ctl update
```

## Features

- Cross-platform builds (Linux/macOS/Windows, amd64/arm64)
- Self-update from GitHub Releases
- Taskfile-based build automation
- GitHub Pages documentation

## Template System

This repo serves as a reusable template for Go binary projects. To create your own:

1. Fork or copy this repo
2. Update vars in `Taskfile.yml`:
   ```yaml
   GITHUB_USER: your-username
   GITHUB_REPO: your-repo
   BINARIES: your-binary
   ```
3. Update `internal/version/version.go` with your repo
4. Add your binary in `cmd/your-binary/`
5. Tag `v0.1.0` to trigger first release

See [MAINTAINERS.md](docs/MAINTAINERS.md) for full details.

## Development

```sh
task build:local        # build for host platform
task build:all          # build all platforms
task test:all           # run all tests
task run:fakeprinter    # run fake printer locally
task run:x1ctl -- status --ip 192.168.1.x --access-code XXX
```

## License

MIT
