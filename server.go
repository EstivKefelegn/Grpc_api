package main

import (
	"context"
	"fmt"
	pb "gRPCApi/golang/proto/gen"
	farewellpb "gRPCApi/golang/proto/gen/farewell"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculatorServer
	pb.UnimplementedGreeterServer
	pb.BidFarewellServer
	farewellpb.AuefWiedersehenServer
}

// mustEmbedUnimplementedGreeterServer implements mainapipb.GreeterServer.
func (c *server) mustEmbedUnimplementedGreeterServer() {
	panic("unimplemented")
}

func (c *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum: ", sum)

	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func (c *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("hello %v nice to see you here", req.Name)
	log.Println("Message: ", message)

	
	return &pb.HelloResponse{
		Message: message,
	}, nil
}

func (c *server) BidGoodBye(ctx context.Context, req *farewellpb.GoodByeRequest) (*farewellpb.GoodByeResponse, error){
	message := fmt.Sprintf("Hello %v. nice to receive request form you. Farewell my friend", req.Name)
	log.Println("Message: ", message)

	return &farewellpb.GoodByeResponse{
		Message: message,
	}, nil
}	

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)

	cert := "cert.pem"
	key := "key.pem"

	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalln("Failed to load credentials: ", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculatorServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})
	pb.RegisterBidFarewellServer(grpcServer, &server{})
	farewellpb.RegisterAuefWiedersehenServer(grpcServer, &server{})

	log.Println("Listening on port: 50051 ...")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
