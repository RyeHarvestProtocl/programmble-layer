package server

import (
	"log"

	"github.com/RyeHarvestProtocol/programmable-layer/programmableLayer"
)

// Initialize sets up the RPC server but does not start it
func Initialize(addr string) {
	log.Println("starting RPC server...")
}

// StartServer starts the RPC server
func StartServer(addr string) {
	go func() {
		programmableLayer.StartServer(addr)
	}()
}

// StopServer stops the RPC server with a timeout
func StopServer() {
	// todo
	log.Println("shuting down RPC server...")
}

func CreateServer(modes ...string) {
	// mode := "" // Default mode
	// if len(modes) > 0 {
	// 	mode = modes[0] // Use the first mode if provided
	// }

	// c := config.New(mode, "")
	// dbInstance := models.New(c)

	// create rpc server
	Initialize(":50051")
	StartServer(":50051")
}
