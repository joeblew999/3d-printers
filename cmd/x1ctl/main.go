package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joeblew999/3d-printers/internal/printer"
	"github.com/joeblew999/3d-printers/internal/printer/x1"
	"github.com/joeblew999/3d-printers/internal/version"
)

type options struct {
	ip          string
	accessCode  string
	insecureTLS bool
	timeout     time.Duration
	showVersion bool
	mode        string
	message     string
}

func main() {
	opts := parseFlags()
	if opts.showVersion {
		fmt.Println(version.Version)
		return
	}

	if opts.ip == "" || opts.accessCode == "" {
		flag.Usage()
		os.Exit(2)
	}

	ctx, cancel := context.WithTimeout(context.Background(), opts.timeout)
	defer cancel()

	client, err := x1.Connect(ctx, printer.Options{
		IP:         opts.ip,
		AccessCode: opts.accessCode,
		Insecure:   opts.insecureTLS,
		Timeout:    opts.timeout,
	})
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer client.Close()

	switch opts.mode {
	case "read":
		log.Printf("connected to %s; waiting for first message...", opts.ip)
		msg, err := client.ReadRaw(ctx)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		fmt.Printf("printer said:\n%s\n", string(msg))
	case "status":
		log.Printf("connected to %s; waiting for first message (status)...", opts.ip)
		msg, err := client.ReadRaw(ctx)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		fmt.Printf("printer said:\n%s\n", string(msg))
		if ver := extractVersion(msg); ver != "" {
			fmt.Printf("detected firmware/version: %s\n", ver)
		}
	case "echo":
		log.Printf("connected to %s; sending echo payload...", opts.ip)
		payload := map[string]any{
			"cmd": "echo",
			"msg": opts.message,
		}
		if err := client.SendJSON(ctx, payload); err != nil {
			log.Fatalf("send: %v", err)
		}
		resp, err := client.ReadRaw(ctx)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		fmt.Printf("printer replied:\n%s\n", string(resp))
	default:
		log.Fatalf("unknown mode: %s (use read or echo)", opts.mode)
	}
}

func parseFlags() options {
	var opts options

	flag.StringVar(&opts.ip, "ip", "", "Printer LAN IP address")
	flag.StringVar(&opts.accessCode, "access-code", "", "Printer LAN access code (from device screen)")
	flag.BoolVar(&opts.insecureTLS, "insecure", true, "Allow self-signed TLS from printer")
	flag.DurationVar(&opts.timeout, "timeout", 15*time.Second, "Dial/read timeout")
	flag.BoolVar(&opts.showVersion, "version", false, "Print version and exit")
	flag.StringVar(&opts.mode, "mode", "read", "Mode: read (wait for first message), status (read + try to detect firmware), or echo (send a test payload)")
	flag.StringVar(&opts.message, "message", "hello from x1ctl", "Echo message to send when mode=echo")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -ip <printer-ip> -access-code <code> [-mode read|status|echo]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	return opts
}

// extractVersion walks a JSON blob looking for version-ish fields.
func extractVersion(data []byte) string {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return ""
	}
	var hits []string
	var walk func(any)
	walk = func(val any) {
		switch t := val.(type) {
		case map[string]any:
			for k, v2 := range t {
				lk := strings.ToLower(k)
				if s, ok := v2.(string); ok && (strings.Contains(lk, "firmware") || strings.Contains(lk, "version") || lk == "ver") {
					hits = append(hits, fmt.Sprintf("%s=%s", k, s))
				}
				walk(v2)
			}
		case []any:
			for _, v2 := range t {
				walk(v2)
			}
		}
	}
	walk(v)
	if len(hits) == 0 {
		return ""
	}
	return strings.Join(hits, "; ")
}
