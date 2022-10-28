package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"memcach/pkg/cache"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := cache.NewCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(c.Set(ctx, &cache.Item{Key: "123", Value: "dfgjkng"}))

	fmt.Println("-----")

	fmt.Println(c.Get(ctx, &cache.Key{Key: "123"}))
}
