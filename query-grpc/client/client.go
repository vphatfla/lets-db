package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/vphatfla/lets-db/query-grpc/query"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var (
	serverAddr = flag.String("addr", "localhost:50051", "Server addr in format host:port")
)

func sendQuery(client pb.QueryServiceClient, message string) {
	log.Printf("Sending and printing result, message = %s \n", message)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	q := &pb.Query{Message: message}
		
	res, err := client.ProcessQuery(ctx, q)
	
	if err != nil {
		log.Fatalf("Error processing query %v \n", err)
	}

	log.Printf("Result = %v \n", res)
}
func main() {
	flag.Parse()
	
	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v \n", err )
	}
	defer conn.Close()
	
	client := pb.NewQueryServiceClient(conn)

	// print hello world
	sendQuery(client, "Hello World gRPC!!!")
}
