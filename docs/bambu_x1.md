# Bambu Lab X1 (control notes)

## Downloads (x1ctl)
- Latest release: [releases/latest](https://github.com/joeblew999/3d-printers/releases/latest)
- Direct binaries (always newest tag):
  - [Linux amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_amd64)
  - [Linux arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_linux_arm64)
  - [macOS amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_amd64)
  - [macOS arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_darwin_arm64)
  - [Windows amd64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_amd64.exe)
  - [Windows arm64](https://github.com/joeblew999/3d-printers/releases/latest/download/x1ctl_windows_arm64.exe)

## What it is (quick)
- CoreXY FDM with enclosure, input-shaping, active bed leveling.
- X1 Carbon adds hardened nozzle + higher temp hardware.
- Build ~256 × 256 × 256 mm, 0.4 mm nozzle (swappable), bed to ~110 °C.
- AMS (4-bay) for multi-material; accepts standard G-code (with Bambu metadata) and `.3mf` projects.

## Control paths (how to talk to it)
- **LAN Mode (local, preferred for automation):**
  - Enable LAN Mode on the printer; note IP and access code (screen).
  - Open TLS WebSocket to `wss://<ip>:8883` (self-signed cert; use `InsecureSkipVerify` in dev).
  - Authenticate by sending JSON login with the access code; exchange JSON messages thereafter.
  - Discovery is usually via printer UDP broadcast; you can also set static IP and skip discovery.
- **Cloud:** Bambu account via Bambu Studio; traffic goes through vendor cloud.
- **Offline file:** microSD with G-code; no network needed.

## Message/protocol notes (LAN)
- Transport: MQTT-like JSON over TLS WebSocket on 8883 (self-signed cert).
- Auth: send `{"cmd":"login","password":"<access code>"}`.
- Typical flows: keepalives, status telemetry, start/stop/pause, and job upload (G-code frames/metadata). Message schema can change with firmware—confirm against your version.
- Keep your network trusted; do not expose 8883 to untrusted clients.

## Minimal Go sketch (adapt to your firmware)
```go
d := websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
conn, _, _ := d.Dial("wss://192.168.1.50:8883", nil)
defer conn.Close()
conn.WriteJSON(map[string]any{"cmd": "login", "password": "ABCD"})
_, msg, _ := conn.ReadMessage()
log.Printf("printer said: %s", msg)
```
Use the OrcaSlicer LAN sender as the reference for full flows (auth, topics, keepalives, job upload).

## References
- OrcaSlicer (LAN implementation, MIT): https://github.com/SoftFever/OrcaSlicer
- Bambu Studio (vendor source): https://github.com/bambulab/BambuStudio
- Protocol notes (community): https://github.com/bambulab/BambuStudio/blob/main/resources/protocol/README.md
- Home Assistant integration (practical LAN client): https://github.com/greghesp/ha-bambulab

## What to verify when scripting
- Firmware version vs. protocol: fields/commands can change—capture traffic on your firmware.
- LAN Mode enabled and access code correct; same subnet; firewall allows 8883.
- Use `-insecure` only on trusted networks; for production, pin the printer cert if possible.
