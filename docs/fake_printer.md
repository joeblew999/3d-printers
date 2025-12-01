# Demo Printer (Placeholder)

This is a placeholder printer profile to exercise the docs layout and downloads flow.

## Downloads (x1ctl)
- Latest release: https://github.com/joeblew999/3d-printers/releases/latest
- Direct binaries (always newest tag):
  - Linux amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_amd64
  - Linux arm64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_arm64
  - macOS amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_amd64
  - macOS arm64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_arm64
  - Windows amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_amd64.exe
  - Windows arm64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_arm64.exe

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
3) Run:
   ```sh
   x1ctl -ip <ip> -access-code <code> -version   # check build/tag
   x1ctl -ip <ip> -access-code <code>            # connect (replace with real target when available)
   ```

## Notes
- Replace this file with actual details when you add another printer; keep links updated in `docs/index.html` and `docs/README.md`.
