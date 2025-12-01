# Docs maintenance (for contributors)

- `docs/index.html` is the GitHub Pages landing page with download links to `releases/latest`.
- `docs/bambu_x1.md` holds the Bambu Lab X1 notes (LAN setup, references).
- `docs/fake_printer.md` is a placeholder/template for additional printers; keep docs/index.html links in sync.
- The landing shows a release badge (`releases/latest`); binaries embed the version from `git describe` via Taskfile `-ldflags`.
- If you add more pages, link them from `index.html` and keep them pure static (no build step).
- Pages deploys from `docs/` via `.github/workflows/pages.yml` on pushes to `main`/`master`.
