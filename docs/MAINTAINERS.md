# Maintainers / Template Guide

This project uses a reusable pattern for Go binaries with:
- Cross-platform builds (Linux/macOS/Windows, amd64/arm64)
- Self-update from GitHub Releases
- GitHub Pages documentation
- Taskfile-based automation

## Forking for a New Project

1. Fork or copy this repo
2. Update `Taskfile.yml` vars:
   ```yaml
   vars:
     GITHUB_USER: your-username
     GITHUB_REPO: your-repo-name
     BINARIES: your-binary-name  # space-separated if multiple
   ```
3. Update `internal/version/version.go`:
   ```go
   var (
       GitHubUser = "your-username"
       GitHubRepo = "your-repo-name"
   )
   ```
4. Add your binary in `cmd/your-binary-name/main.go`
5. Push and tag `v0.1.0` to trigger first release

## Project Structure

```
cmd/
  x1ctl/           # Cobra-based CLI
  fakeprinter/     # flag-based CLI
internal/
  version/         # Self-update logic (reusable)
docs/
  index.md         # Landing page (GitHub Pages)
  *_user.md        # User guides
  *_tech.md        # Technical docs
.github/
  workflows/
    ci.yml         # Build + release on tags
```

## Self-Update Mechanism

The `internal/version/` package provides:
- `version.Info()` - returns version string
- `version.CheckUpdate()` - queries GitHub API for latest release
- `version.SelfUpdate(binaryName)` - downloads and replaces the running binary

Build-time version injection via Taskfile:
```yaml
LD_FLAGS: "-X {{.GO_MODULE}}/internal/version.Version={{.VERSION}}"
```

## CI/CD

GitHub Actions (`.github/workflows/ci.yml`):
- **build** job: runs on push, calls `task build:all`
- **release** job: runs on `v*` tags, uploads to GitHub Releases

Key task used by CI: `task build:all`

## GitHub Pages

Pages serve from `/docs` on main branch (native GitHub markdown rendering).

Setup tasks:
```sh
task docs:pages:setup   # configure Pages via gh API
task docs:pages:status  # check configuration
task docs:serve         # local preview (requires static-server)
```

## Common Tasks

```sh
task build:local        # build for current platform
task build:all          # build all platforms
task install:local      # build + install to /usr/local/bin
task install:remote     # install latest from GitHub Releases
task release:create -- v0.2.0  # tag + push (triggers release)
task test:all           # run all tests
```

## Adding a New Binary

1. Create `cmd/newbinary/main.go`
2. Add to `BINARIES` in `Taskfile.yml`:
   ```yaml
   BINARIES: x1ctl fakeprinter newbinary
   ```
3. For self-update support, add version/update flags or subcommands (see existing binaries for patterns)

## Doc Conventions

- `*_user.md` - end-user guides (downloads, quick start)
- `*_tech.md` - technical/protocol details
- Keep docs pure markdown (no build step)
- Link from `index.md`
