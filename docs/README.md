# 3d-printers docs (user view)

Welcome! Grab the CLI and printer guides here.

- Downloads (latest release): https://joeblew999.github.io/3d-printers/  
  Binaries always point to the newest tag; check your build with `x1ctl -version`.
- Release page (assets + notes): https://github.com/joeblew999/3d-printers/releases/latest

## Printer guides
- Bambu Lab X1: `docs/bambu_x1_user.md` (user), `docs/bambu_x1_tech.md` (technical)
- Demo printer: `docs/fake_printer_user.md` (user), `docs/fake_printer_tech.md` (technical)
- Contribute/fix docs: https://github.com/joeblew999/3d-printers/tree/main/docs · Issues: https://github.com/joeblew999/3d-printers/issues

## Quick start
1) Download the binary for your OS from the Downloads page.  
2) On the printer, enable LAN Mode; note IP + access code.  
3) Run `./x1ctl -ip <printer-ip> -access-code <code> -version` to confirm the build, then run without `-version` to connect.
