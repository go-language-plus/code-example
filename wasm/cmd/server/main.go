// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":2345", "listen address")
	dir    = flag.String("dir", "../../assets", "directory to serve")
)

// go run cmd/server/main.go -dir assets/
func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
