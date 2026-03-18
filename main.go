package main

import (
	"flag"

	"github.com/kuche1/cloud-note/client"
	"github.com/kuche1/cloud-note/server"
)

func main() {
	runServer := flag.Bool("run-server", false, "Run server")
	serverAddress := flag.String("server-address", "localhost:4242", "Address to run server on")
	flag.Parse()

	if *runServer {
		server.Main(*serverAddress)
	} else {
		client.Main()
	}
}
