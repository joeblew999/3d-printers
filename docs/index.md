# x1ctl downloads

Latest binaries are attached to GitHub releases. These links always point to the newest tagged release (`v*`):

- Linux amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_amd64
- Linux arm64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_arm64
- macOS amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_amd64
- macOS arm64 (Apple Silicon): https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_arm64
- Windows amd64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_amd64.exe
- Windows arm64: https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_arm64.exe

## Usage
```sh
./x1ctl -ip 192.168.1.50 -access-code ABCD
```

## How releases are built
- CI builds the matrix on every push/PR and on tags (`v*`), then attaches the binaries to the release.
- These links use `releases/latest`, so tagging a new version automatically updates the downloads above.
