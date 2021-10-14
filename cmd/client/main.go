/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/wwestgarth/remote-geometry/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
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
