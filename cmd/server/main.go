package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"memcach/pkg/cache"
	"memcach/pkg/server"
	"net"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.CacheServer{}
	cache.RegisterCacheServer(s, &srv)

	log.Printf("Starting gRPC listener on port " + os.Getenv("PORT"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
