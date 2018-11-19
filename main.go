package main

import (
	"context"
	"github.com/vllry/k2/api/0.0"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewKube2Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.CreateContainer(ctx, &api.CreateContainerRequest{Image: "nginx", ImageTag: "latest"})
	if err != nil {
		log.Fatalf("could not launch: %v", err)
	}
	log.Printf("Status: %b", r.Success)
}