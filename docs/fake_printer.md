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
- A fake printer service to test x1ctl and the release pipeline.
- Behavior: accepts TLS websocket, sends a hello JSON, then echoes any JSON you send. Access code is ignored (for testing only).

## Setup checklist (for the mock)
1) Start the mock printer: `task fakeprinter` (TLS websocket on `wss://localhost:8883` with self-signed cert).
   - Or run the binary directly, e.g. `./fakeprinter_darwin_arm64 -addr :8883`
2) Connect with x1ctl:
   ```sh
   x1ctl -ip localhost -access-code any -insecure -version   # check build/tag
   x1ctl -ip localhost -access-code any -insecure             # connect to mock
   ```

## Fake printer server (local)
- Uses a self-signed cert; `-insecure` on x1ctl should stay true for this mock.
- Sends an initial message like `{"hello":"fake-printer","ts":<unix>}` then echoes any JSON back with a timestamp.
- Use Ctrl+C to stop the mock server.

## Notes
- Replace this file with actual details when you add another printer; keep links updated in `docs/index.html` and `docs/README.md`.
