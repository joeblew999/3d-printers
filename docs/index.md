# 3d-printers

Go binaries with self-update from GitHub Releases.

## Binaries

| Binary | Description |
|--------|-------------|
| **x1ctl** | CLI for Bambu Lab X1 printers |
| **fakeprinter** | Demo/test printer server |

## Downloads

[Latest Release](https://github.com/joeblew999/3d-printers/releases/latest) | [All Releases](https://github.com/joeblew999/3d-printers/releases)

| Binary | Linux | macOS | Windows |
|--------|-------|-------|---------|
| x1ctl | [amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_amd64) / [arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_arm64) | [amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_amd64) / [arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_arm64) | [amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_amd64.exe) / [arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_arm64.exe) |
| fakeprinter | [amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_linux_amd64) / [arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_linux_arm64) | [amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_darwin_amd64) / [arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_darwin_arm64) | [amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_windows_amd64.exe) / [arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/fakeprinter_windows_arm64.exe) |

## Self-Update

All binaries support self-update from GitHub Releases:

```sh
# x1ctl (Cobra subcommands)
x1ctl version          # show version
x1ctl version --check  # check for updates
x1ctl update           # download and replace binary

# fakeprinter (flag-based)
fakeprinter --version       # show version
fakeprinter --check-update  # check for updates
fakeprinter --update        # download and replace binary
```

## Documentation

- [Bambu X1 User Guide](bambu_x1_user.md) | [Technical](bambu_x1_tech.md)
- [Fake Printer Guide](fake_printer_user.md) | [Technical](fake_printer_tech.md)

## For Developers

- [Contributing / Forking this Template](MAINTAINERS.md)
- [Source on GitHub](https://github.com/joeblew999/3d-printers)
- [File an Issue](https://github.com/joeblew999/3d-printers/issues)
