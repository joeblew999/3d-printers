# 3d-printers
Utilities and notes for Bambu Lab X1 (see `docs/bambu_x1.md`). The `x1ctl` CLI is a small Go tool to experiment with LAN Mode.

## Requirements
- Go 1.22+
- Task (https://taskfile.dev) for builds: `brew install go-task/tap/go-task` or see their install docs.

## Getting started
```sh
# build all target binaries (linux/windows/darwin × amd64/arm64)
task build:all

# run tests
task test

# connect to a printer (example)
./dist/x1ctl_darwin_arm64 -ip 192.168.1.50 -access-code ABCD
# check version embedded at build time
./dist/x1ctl_darwin_arm64 -version
```

## Project layout
- `cmd/x1ctl`: CLI entrypoint.
- `internal/lan`: minimal LAN client for the printer (TLS WebSocket + login).
- `docs/`: GitHub Pages content with download links and docs (`bambu_x1.md`).
- `internal/version`: version string injected via `-ldflags` from `git describe`.
- `Taskfile.yml`: build matrix and CI helper tasks.

## Release flow
- CI (`.github/workflows/ci.yml`) runs tests, builds the matrix, and uploads artifacts on every push/PR.
- Tagging `v*` triggers a release job that rebuilds via `task build:all` and attaches the binaries to the GitHub release.
- Manual release steps if desired:
  ```sh
  git tag v0.1.0
  git push origin v0.1.0
  ```

## Downloads (GitHub Pages)
- Download page (GitHub Pages): https://joeblew999.github.io/3d-printers/ (links to the six binaries)
- Direct links resolve to the most recent tag (`releases/latest`); see `docs/index.md` for the matrix and usage snippet.

## GitHub Pages setup
- Pages workflow lives at `.github/workflows/pages.yml` and deploys `docs/` on push to `main`/`master`.
- Enable Pages here: https://github.com/joeblew999/3d-printers/settings/pages → set Source to “GitHub Actions”.
- The public landing page (with download links) will be at `https://joeblew999.github.io/3d-printers/` after the first successful deploy.

## Development tips
- `task ci` runs tests and builds the matrix locally.
- `task clean` removes `dist/`.
- Keep Go code formatted with `gofmt` (run automatically by your editor or `gofmt -w`).
