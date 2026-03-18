package main

import (
	"flag"

	"github.com/kuche1/cloud-note/client"
	"github.com/kuche1/cloud-note/server"
)

func main() {
	runServer := flag.Bool("run-server", false, "Run server")
	serverAddress := flag.String("server-address", "localhost:4242", "Address to run server on")
	serverFilesystemStorage := flag.String("server-filesystem-storage", "./server-storage", "Where to store the server files")
	flag.Parse()

	if *runServer {
		server.Main(*serverAddress, *serverFilesystemStorage)
	} else {
		client.Main()
	}
}
