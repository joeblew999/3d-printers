# Bambu Lab X1 quick notes

## What it is
- FDM printer (CoreXY) with enclosed chamber, input-shaping, and active bed leveling.
- X1 Carbon variant adds hardened nozzle, higher temp hardware, and better cooling; otherwise similar workflow.

## Core specs (operational highlights)
- Build volume: ~256 × 256 × 256 mm.
- Nozzle: 0.4 mm stock (swapable), up to ~300 °C; bed up to ~110 °C.
- Multi-material: optional AMS (4-bay filament feeder).
- File formats: standard G-code with Bambu-flavored extensions/metadata; `.3mf` project files from slicers.

## How to talk to it
- **Official path (supported):** Bambu Studio slicer (or OrcaSlicer) → send over LAN or cloud.
  - On the printer, enable **LAN Mode** (to allow local control without cloud). Note the IP shown on-screen.
  - In Studio/Orca, add device → “LAN only” → enter the printer’s IP and access code shown on the printer.
  - Send prints or start/stop jobs directly over LAN; camera/telemetry also flow through the slicer UI.
- **Cloud path:** Same slicer, but sign in to Bambu account; prints route via Bambu’s cloud services.
- **Local file:** Export G-code to microSD and print from the printer menu (bypasses network).

## Useful tools (open-source or community)
- **OrcaSlicer** (MIT-licensed fork of Bambu Studio) – good reference for how LAN control is implemented and for extra tuning features.
- **Bambu Studio** (official slicer, source available) – matches vendor UX, good for comparison.
- Community protocol notes exist for LAN control (see links below); the protocol is not fully published by Bambu, so rely on up-to-date community findings.
- Common workflows for automation integrate via OrcaSlicer’s LAN code paths or by emitting G-code to the printer’s SD/LAN sender.

## Safety and constraints
- The printer firmware and many cloud pieces are proprietary; there is no official public API spec.
- Firmware updates can change LAN behaviors/protocol details; verify against the printer’s current firmware version.
- Treat LAN access as trusted-only: avoid exposing the printer to untrusted networks; use VLAN/wifi isolation if needed.

## Quick checklist to get connected on LAN
1) On printer: enable **LAN Mode**; note IP + access code.  
2) On PC: install Bambu Studio (or OrcaSlicer).  
3) Add device → choose LAN → enter IP + code → test connection (see status indicator).  
4) Send a small test print; confirm monitoring/controls work.  
5) If connection fails: ensure same subnet, no firewall blocking slicer’s ports, and printer still in LAN Mode.

## Go LAN control sketch (LAN Mode)
- The LAN interface uses MQTT-like traffic over a TLS WebSocket on port `8883` with a self-signed cert; discovery typically happens via UDP broadcast from the printer.
- Minimal flow: discover IP (or set static), get the on-screen access code, open a TLS WebSocket to `wss://<printer-ip>:8883` using that code in the auth payload, then publish/subscribe JSON messages (status, commands, g-code frames). Message details can change with firmware; confirm against your printer.
- Example outline (using `github.com/gorilla/websocket`):
```go
package main

import (
	"crypto/tls"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	ip := "192.168.1.50" // printer IP from LAN Mode screen
	accessCode := "ABCD" // printer's LAN access code

	d := websocket.Dialer{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // printer uses self-signed cert
	}

	conn, _, err := d.Dial("wss://"+ip+":8883", nil)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	auth := map[string]any{
		"cmd":  "login",
		"password": accessCode,
	}
	if err := conn.WriteJSON(auth); err != nil {
		log.Fatalf("auth send: %v", err)
	}

	// Read responses and publish commands per protocol notes...
}
```
- For a fuller reference, review OrcaSlicer’s LAN sender code (auth handshake, keepalives, message schemas) and mirror that behavior in Go.

## References and links
- OrcaSlicer (MIT): https://github.com/SoftFever/OrcaSlicer
- Bambu Studio source: https://github.com/bambulab/BambuStudio
- Home Assistant LAN integration (good protocol reference): https://github.com/greghesp/ha-bambulab
- Community LAN protocol notes: https://github.com/bambulab/BambuStudio/blob/main/resources/protocol/README.md (check against your firmware)

## What to explore next
- Review OrcaSlicer source to understand LAN request/response flow and message formats.
- Capture traffic while sending a job to confirm ports/payloads used by your firmware version.
- Scripted control can be built by mimicking the slicer’s LAN calls, but budget time for protocol changes between firmware releases.
