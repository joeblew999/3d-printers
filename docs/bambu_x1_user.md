# Bambu Lab X1 (user)

## Downloads (x1ctl)
- Latest release: [releases/latest](https://github.com/joeblew999/3d-printers/releases/latest)
- Direct binaries:
  - [Linux amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_amd64)
  - [Linux arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_arm64)
  - [macOS amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_amd64)
  - [macOS arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_arm64)
  - [Windows amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_amd64.exe)
  - [Windows arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_arm64.exe)

## Quick connect (LAN Mode)
1) On the printer: enable **LAN Mode**, note the IP and access code.
2) On your PC: run x1ctl (self-signed TLS; `-insecure` is on by default):
   ```sh
   x1ctl -ip 192.168.1.50 -access-code ABCD -mode status   # read first message and try to show firmware
   x1ctl -ip 192.168.1.50 -access-code ABCD -mode echo -message "hi"   # send/receive JSON
   ```
3) If it fails: ensure same subnet, port 8883 open, LAN Mode still on, access code correct.

## Notes
- `-version` prints the embedded build version.
- Keep LAN access on trusted networks; the printer uses a self-signed cert.
- For protocol details and scripting, see the technical doc: `bambu_x1_tech.md`.
