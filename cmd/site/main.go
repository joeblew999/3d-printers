package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	dir := flag.String("dir", "docs", "directory to serve")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))
	log.Printf("serving %s on %s", *dir, *addr)
	log.Fatal(http.ListenAndServe(*addr, fs))
}
