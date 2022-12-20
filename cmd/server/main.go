package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"memcach/pkg/cache"
	"memcach/pkg/inmemory"
	"memcach/pkg/memcache"
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
	mc, err := memcache.New(os.Getenv("TELNET-PORT"))
	if err != nil {
		panic(err)
	}
	defer mc.Close()
	ns := inmemory.NewStorage(mc)
	srv := server.NewCacheServer(ns)
	s := grpc.NewServer()
	cache.RegisterCacheServer(s, srv)

	log.Printf("Starting gRPC listener on port " + os.Getenv("PORT"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
