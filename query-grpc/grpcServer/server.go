package grpcServer

import (
	"context"
	"log"

	pb "github.com/vphatfla/lets-db/query-grpc/query"
)

type QuerySeviceServer struct {
	pb.UnimplementedQueryServiceServer
}

func (s *QuerySeviceServer) ProcessQuery (_ context.Context, query *pb.Query) (*pb.Result, error) {
	log.Printf("Received query with message %v \n", query.Message)

	return &pb.Result{Message: query.Message}, nil
}

func NewServer() *QuerySeviceServer {
	s := &QuerySeviceServer{}
	return s
}

/* func main() {
	log.Println("Starting the grpc server ")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterQueryServiceServer(grpcServer, newServer())

	log.Printf("GRPC Serve and Listen on port %d", *port)
	grpcServer.Serve(lis)
} */
