package programmableLayer

import (
	"context"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

// TestSayHello tests the SayHello RPC method.
func TestSayHello(t *testing.T) {
	// Set up a listener for the gRPC server.
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server and register the Greeter service.
	grpcServer := grpc.NewServer()
	RegisterGreeterServer(grpcServer, &Server{})

	// Use a WaitGroup to wait for the server to start.
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the server in a goroutine.
	go func() {
		defer wg.Done() // Signal that the goroutine has completed
		if err := grpcServer.Serve(lis); err != nil {
			// Use a log function here to avoid calling Fatalf from a goroutine
			t.Logf("failed to serve: %v", err)
		}
	}()

	// Ensure the server has started before continuing.
	defer func() {
		grpcServer.Stop()
		wg.Wait() // Wait for the server goroutine to finish
	}()

	// Create a gRPC client connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new Greeter client.
	client := NewGreeterClient(conn)

	// Call the SayHello method.
	req := &HelloRequest{Name: "World"}
	resp, err := client.SayHello(context.Background(), req)

	// Check for errors and validate the response.
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}

	expectedMessage := "Hello World"
	assert.Equal(t, expectedMessage, resp.Message)
}
