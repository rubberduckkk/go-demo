package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rubberduckkk/go-demo/grpc/metadata/pb"
)

var (
	flagPort = flag.Int("port", 8789, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8789))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDemoServiceServer(grpcServer, newService())
	grpcServer.Serve(lis)
}
