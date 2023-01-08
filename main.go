package main

import (
	"fmt"
	"flag"
	"log"
	"time"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/sekuradev/apigolang/sekuraapi/v1"
)

var serverName string
var serverPort int
var certFile string

func send(){

	serverAddress := fmt.Sprintf("%s:%d", serverName, serverPort)

    creds, err := credentials.NewClientTLSFromFile(certFile, serverName)
    if err != nil {
        log.Fatalf("error creating client credentials: %v", err)
    }

	conn, errConn := grpc.Dial(serverAddress, grpc.WithTransportCredentials(creds))
	if errConn != nil {
		log.Fatalf("Connection to server could not be stablished at %s: %v", serverAddress, errConn)
	}
	defer conn.Close()

	client := pb.NewAgentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := &pb.SetAccessRequest{
	}

	_, errClient := client.SetAccess(ctx, data)
	if errClient != nil {
		log.Printf("error connecting: %v", errClient)
	}
}

func main() {
	flag.StringVar(&serverName, "serverName", "sekura.localhost", "Server address.")
	flag.IntVar(&serverPort, "serverPort", 50051, "Server port.")
    flag.StringVar(&certFile, "cert", "", "Server certificate file.")
    flag.Parse()

	if certFile == "" {
        log.Fatalln("flag --cert is required")
    }

	fmt.Printf("Connecting to %s:%d\n", serverName, serverPort)
	for true{
		log.Println("Sending data.")
		send()
		time.Sleep(10 * time.Second)
	}
}
