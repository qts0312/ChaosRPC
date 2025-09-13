package main

import (
	"context"
	pb "github.com/qts0312/ChaosRPC/examples/toy/proto"
	chaosgrpc "github.com/qts0312/ChaosRPC/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ToyServer struct {
	pb.UnimplementedToyServiceServer
}

func (s *ToyServer) Handshake(ctx context.Context, req *pb.HandshakeRequest) (*pb.HandshakeResponse, error) {
	return &pb.HandshakeResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	chaosgrpc.Init()

	conn, err := net.Listen("tcp", ":50051")
	defer conn.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterToyServiceServer(s, &ToyServer{})
	log.Printf("server listening at %v", conn.Addr())
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
