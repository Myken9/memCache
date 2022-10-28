package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"memcach/pkg/cache"
	"memcach/pkg/inmemory"
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
	st := inmemory.NewStorage(make(map[string]string))
	srv := server.NewCacheServer(st)
	cache.RegisterCacheServer(s, srv)

	log.Printf("Starting gRPC listener on port " + os.Getenv("PORT"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
