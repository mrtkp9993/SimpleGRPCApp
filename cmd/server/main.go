package main

import (
	"SimpleGRPCApp/numberAPI"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	lis, err := net.Listen(viper.GetString("server.network"), viper.GetString("server.port"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := numberAPI.Server{}

	creds, err := credentials.NewServerTLSFromFile(viper.GetString("cert.certfile"), viper.GetString("cert.keyfile"))
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
