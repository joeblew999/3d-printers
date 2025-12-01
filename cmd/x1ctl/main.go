package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joeblew999/3d-printers/internal/lan"
	"github.com/joeblew999/3d-printers/internal/version"
)

type options struct {
	ip          string
	accessCode  string
	insecureTLS bool
	timeout     time.Duration
	showVersion bool
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

	client, err := lan.Dial(ctx, opts.ip, opts.accessCode, opts.insecureTLS)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer client.Close()

	log.Printf("connected to %s; waiting for first message...", opts.ip)

	// Read the first response to confirm the connection.
	msg, err := client.ReadRaw(ctx)
	if err != nil {
		log.Fatalf("read: %v", err)
	}
	fmt.Printf("printer said:\n%s\n", string(msg))
}

func parseFlags() options {
	var opts options

	flag.StringVar(&opts.ip, "ip", "", "Printer LAN IP address")
	flag.StringVar(&opts.accessCode, "access-code", "", "Printer LAN access code (from device screen)")
	flag.BoolVar(&opts.insecureTLS, "insecure", true, "Allow self-signed TLS from printer")
	flag.DurationVar(&opts.timeout, "timeout", 15*time.Second, "Dial/read timeout")
	flag.BoolVar(&opts.showVersion, "version", false, "Print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -ip <printer-ip> -access-code <code>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	return opts
}
