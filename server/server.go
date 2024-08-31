package server

import (
	"fmt"
	"log"

	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/RyeHarvestProtocol/programmable-layer/programmableLayer"
)

// Initialize sets up the RPC server but does not start it
func Initialize(addr string) {
	log.Println("starting RPC server...")
}

// StartServer starts the RPC server
func StartServer(addr string, c *config.Config) {
	go func() {
		programmableLayer.StartServer(addr, c)
	}()
}

// StopServer stops the RPC server with a timeout
func StopServer() {
	// todo
	log.Println("shuting down RPC server...")
}

func CreateServer(modes ...string) {
	mode := "" // Default mode
	if len(modes) > 0 {
		mode = modes[0] // Use the first mode if provided
	}

	c := config.New(mode, "")

	fmt.Println("this is cccc: ", c)
	// dbInstance := models.New(c)

	// create rpc server
	Initialize(":50051")
	StartServer(":50051", c)
}
