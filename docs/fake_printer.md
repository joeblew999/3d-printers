# Demo Printer (Placeholder)

This is a placeholder printer profile to exercise the docs layout and downloads flow.

## What it is
- A fake FDM printer profile for testing documentation and release pipeline.
- Use this page as a template for adding real printers.

## LAN/control sketch
- Control path: same CLI (`x1ctl`) for LAN tests; substitute a test IP/access code when mocking.
- For real printers, document discovery, auth, and supported commands here.

## Specs (stub)
- Build volume: TBD
- Nozzle: TBD
- Materials: TBD

## Setup checklist (stub)
1) Enable LAN mode (or equivalent) on the device.
2) Note IP and access code.
3) Run `x1ctl -ip <ip> -access-code <code> -version` to confirm build; then run without `-version` to connect.

## Notes
- Replace this file with actual details when you add another printer; keep links updated in `docs/index.html` and `docs/README.md`.
