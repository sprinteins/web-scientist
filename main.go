package main

import (
	"flag"
	"fmt"

	"github.com/trusz/web-scientist/server"
)

func main() {
	var host = flag.String("host", "localhost", "The Host")
	var port = flag.String("port", "7654", "The port")
	flag.Parse()
	fmt.Printf("listening on http://%s:%s\n", *host, *port)
	// server(*host, *port)
	fmt.Printf("server started")
	// server.Start(*host, *port)
	var scientist = server.New(*host, *port)
	scientist.Start()
}
