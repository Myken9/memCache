package main

import (
	"google.golang.org/grpc"
	"log"
	"memcach/pkg/cache"
	"memcach/pkg/server"
	"net"
)

const (
	port = ":50051"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.CacheServer{}
	cache.RegisterCacheServer(s, &srv)

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
