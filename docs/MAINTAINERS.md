# Docs maintenance (for contributors)

- `docs/index.md` is the Markdown landing; CI builds static HTML via `cmd/sitebuild` into `.site/` for Pages.
- Per-printer split: `*_user.md` for quick steps/downloads, `*_tech.md` for protocol/impl details.
- Release artifacts now include `x1ctl_*` and `fakeprinter_*` per OS/arch; Task `build:all` builds both.
- x1ctl CLI modes: read/status/echo (see `cmd/x1ctl` flags).
- The landing shows a release badge (`releases/latest`); binaries embed the version from `git describe` via Taskfile `-ldflags`.
- If you add more pages, link them from `index.html` and keep them pure static (no build step).
- Pages deploys from `docs/` via `.github/workflows/pages.yml` on pushes to `main`/`master`.
