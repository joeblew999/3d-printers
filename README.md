# 3d-printers
# 3d-printers
CLI and docs for Bambu Lab X1 and a demo (fake) printer.

- Docs & downloads: https://joeblew999.github.io/3d-printers/
- Printer pages: user/tech docs per printer (see `docs/*_user.md`, `docs/*_tech.md`).
- Builds on tags (`v*`) produce binaries for Linux/macOS/Windows (amd64/arm64); the docs links always point to the latest release.
- Local preview: `task serve` (docs) and `task fakeprinter` (mock server). Build all binaries: `task build:all`.
- x1ctl modes: `-mode read` (first message), `-mode status` (try to surface firmware fields), `-mode echo` (send/recv JSON).
