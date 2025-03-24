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

func insertQuery(client pb.QueryServiceClient, m string) {
	log.Printf("Sending insert query, message = %s \n", m)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	q := &pb.Query{Message: m}

	if res, err := client.Insert(ctx, q); err != nil {
		log.Println(res.Message)
		log.Fatalln("Error for an insert query ->", err)
	} else {
		log.Println(res.Message)
	}
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

	// send insert query
	insertQuery(client, "insert abc value 123")
}
