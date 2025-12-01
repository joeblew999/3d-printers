# x1ctl downloads

[![latest release](https://img.shields.io/github/v/release/joeblew999/3d-printers?display_name=tag)](https://github.com/joeblew999/3d-printers/releases/latest)

Latest binaries (always point to the most recent tag):
- Linux amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_amd64
- Linux arm64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_arm64
- macOS amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_amd64
- macOS arm64 (Apple Silicon): https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_arm64
- Windows amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_amd64.exe
- Windows arm64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_arm64.exe

## Usage
```sh
./x1ctl -version                  # show build version (from git tag)
./x1ctl -ip 192.168.1.50 -access-code ABCD
```

## Printer docs
- Bambu Lab X1: [bambu_x1.md](bambu_x1.md)
- Demo printer (placeholder): [fake_printer.md](fake_printer.md)

## How releases are built
- CI builds linux/windows/mac binaries for amd64 and arm64 on tags (`v*`) and attaches them to the release.
- The links above use `releases/latest`, so tagging a new version automatically updates downloads.
