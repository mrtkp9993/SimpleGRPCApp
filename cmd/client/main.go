package main

import (
	"SimpleGRPCApp/numberAPI"
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
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
	var conn *grpc.ClientConn

	creds, err := credentials.NewClientTLSFromFile(viper.GetString("cert.certfile"), viper.GetString("cert.servername"))
	if err != nil {
		log.Fatalf("Cannot load TLS file: %s", err)
	}

	conn, err = grpc.Dial(viper.GetString("server.target"), grpc.WithTransportCredentials(creds))
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
