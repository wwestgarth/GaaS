package main

import (
	"context"
	"log"
	"time"

	pb "github.com/wwestgarth/remote-geometry/protos"

	"google.golang.org/grpc"
)

func connect() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRemoteGeometryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateSheetCircle(ctx,
		&pb.CreateSheetCircleRequest{
			Radius: 10,
			Centre: &pb.Vector{X: 0, Y: 0, Z: 0},
			Normal: &pb.Vector{X: 0, Y: 0, Z: 1},
		},
	)
	if err != nil {
		log.Fatalf("could not create body: %v", err)
	}
	log.Printf("BodyID: %s", r.GetId())
}
