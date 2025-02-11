package main

import (
	"flag"
	"fmt"
	"solidcute/muxserver/pkg/server"
)

func main() {
	port := flag.Int("port", 8080, "Host to listen on.")
	host := flag.String("host", "127.0.0.1", "Host to listen on.")

	flag.Parse()

	server.Listen(fmt.Sprintf("%s:%d", *host, *port))
}
