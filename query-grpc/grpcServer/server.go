package grpcServer

import (
	"context"
	"fmt"
	"log"

	"github.com/vphatfla/lets-db/db/parser"
	pb "github.com/vphatfla/lets-db/query-grpc/query"
)

type QuerySeviceServer struct {
	pb.UnimplementedQueryServiceServer
}

func (s *QuerySeviceServer) ProcessQuery (_ context.Context, query *pb.Query) (*pb.Result, error) {
	log.Printf("Received query with message %v \n", query.Message)

	return &pb.Result{Message: query.Message}, nil
}

func (s* QuerySeviceServer) Insert (_ context.Context, query *pb.Query) (*pb.Result, error) {
	log.Printf("Received Insert query with message %v \n", query.Message)

	h, key, value, err := parser.ParseQuery(query.Message)

	if err != nil {
		return nil, fmt.Errorf("Error parsing insert query -> %v \n", err)
	}

	if	err :=	h.Execute(key, value); err != nil {
		return nil, fmt.Errorf("Erroring executing IO insert with query %s -> error = %v", query.Message, err)
	}
	return &pb.Result{Message: "Successed, pending result format"}, nil
}

func NewServer() *QuerySeviceServer {
	s := &QuerySeviceServer{}
	return s
}
