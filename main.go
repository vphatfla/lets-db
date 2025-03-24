package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/vphatfla/lets-db/query-grpc/grpcServer"
	pb "github.com/vphatfla/lets-db/query-grpc/query"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The db server port")
)

func main() {
	log.Println("Starting the DB server")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterQueryServiceServer(s, grpcServer.NewServer())

	log.Printf("DB GRPC Listening on port %d", *port)

	s.Serve(lis)
}
