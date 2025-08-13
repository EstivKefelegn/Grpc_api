package main

import (
	"context"
	pb "gRPCApi/golang/proto/gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (c *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculatorServer(grpcServer, &server{})
	log.Println("Listening on port: 50051 ...")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
