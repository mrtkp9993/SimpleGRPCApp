package main

import (
	"SimpleGRPCApp/numberAPI"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := numberAPI.Server{}
	grpcServer := grpc.NewServer()

	numberAPI.RegisterGetNumberServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
