package main

import (
	"flag"

	"github.com/kuche1/cloud-note/client"
	"github.com/kuche1/cloud-note/server"
)

func main() {
	runServer := flag.Bool("run-server", false, "Run server")
	flag.Parse()

	if *runServer {
		server.Main()
	} else {
		client.Main()
	}
}
