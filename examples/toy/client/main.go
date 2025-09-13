package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/qts0312/ChaosRPC/examples/toy/proto"
	chaosgrpc "github.com/qts0312/ChaosRPC/pkg/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "server address")
)

func main() {
	chaosgrpc.Init()

	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	c := pb.NewToyServiceClient(conn)

	for i := 0; i < 3; i++ {
		ctx := context.Background()
		resp, err := c.Handshake(ctx, &pb.HandshakeRequest{
			Name: "client",
		})
		if err != nil {
			log.Printf("could not handshake: %v", err)
		} else {
			log.Printf("handshake success: %v", resp)
		}
	}
}
