package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/RyeHarvestProtocol/programmable-layer/server"
)

func getPort() string {
	c := config.New("", "")
	port := c.Port

	if port == "" {
		port = "50051"
	}

	return port
}

func main() {
	server.CreateServer()

	fmt.Println("Server is running at localhost:" + getPort())

	// Handling graceful shutdown in a goroutine
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	server.StopServer()
}
