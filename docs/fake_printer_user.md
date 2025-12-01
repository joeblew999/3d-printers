# Demo Printer (user)

This mock printer lets you try x1ctl without real hardware.

## Downloads (fake printer)
- Latest release: [releases/latest](https://github.com/joeblew999/3d-printers/releases/latest)
- Direct binaries:
  - [Linux amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_linux_amd64)
  - [Linux arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_linux_arm64)
  - [macOS amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_darwin_amd64)
  - [macOS arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_darwin_arm64)
  - [Windows amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_windows_amd64.exe)
  - [Windows arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_windows_arm64.exe)

## Quick start
1) Start the mock server: `task fakeprinter` (or run the binary: `./fakeprinter_<os>_<arch> -addr :8883`).
2) Connect with x1ctl:
   ```sh
   x1ctl -ip localhost -access-code any -insecure -mode status   # read first message
   x1ctl -ip localhost -access-code any -insecure -mode echo -message "hi"   # echo test
   ```
3) Stop the mock with Ctrl+C.

## Notes
- Self-signed TLS; keep `-insecure` true for the mock.
- Access code is ignored by the mock (for testing only).
- For implementation details, see `fake_printer_tech.md`.
