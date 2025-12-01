# Demo Printer (Placeholder)

This is a placeholder printer profile to exercise the docs layout and downloads flow.


## Downloads (fake printer mock)
- Latest release: [releases/latest](https://github.com/joeblew999/3d-printers/releases/latest)
- Direct binaries (always newest tag):
  - [Linux amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_linux_amd64)
  - [Linux arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_linux_arm64)
  - [macOS amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_darwin_amd64)
  - [macOS arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_darwin_arm64)
  - [Windows amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_windows_amd64.exe)
  - [Windows arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_windows_arm64.exe)

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

## Setup checklist (for the mock)
1) Start the mock printer: `task fakeprinter` (TLS websocket on `wss://localhost:8883` with self-signed cert).
2) Run:
   ```sh
   x1ctl -ip localhost -access-code any -insecure -version   # check build/tag
   x1ctl -ip localhost -access-code any -insecure             # connect to mock
   ```

## Fake printer server (local)
- Start the mock printer: `task fakeprinter` (listens on `wss://localhost:8883` with self-signed TLS).
- Point `x1ctl` at `localhost` with `-access-code anything` and `-insecure` (default true) to exercise the connection flow.

## Notes
- Replace this file with actual details when you add another printer; keep links updated in `docs/index.html` and `docs/README.md`.
