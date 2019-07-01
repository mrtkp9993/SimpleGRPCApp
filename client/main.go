package main

import (
	"SimpleGRPCApp/numberAPI"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

func main() {
	var conn *grpc.ClientConn

	// Write your Common Name (e.g. server FQDN or YOUR name) as `serverNameOverride`
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("Cannot load TLS file: %s", err)
	}

	conn, err = grpc.Dial(":7777", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Cannot connect: %s", err)
	}
	defer conn.Close()

	c := numberAPI.NewGetNumberClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter number name: ")
		userInput, _ := reader.ReadString('\n')
		request := numberAPI.Request{Name: userInput}
		log.Printf("Your request: %v", request)
		response, err := c.Get(context.Background(), &request)
		if err != nil {
			log.Fatalf("Error when requesting: %s", err)
		}
		log.Printf("Response from server: %v", response.Value)
	}

}
