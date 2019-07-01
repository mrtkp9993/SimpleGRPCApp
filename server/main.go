package main

import (
	"SimpleGRPCApp/numberAPI"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := numberAPI.Server{}

	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("Cannot load TLS file: %s", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	grpcServer := grpc.NewServer(opts...)

	numberAPI.RegisterGetNumberServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
