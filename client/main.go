package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"memcach/pkg/cache"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	conn, err := grpc.Dial(os.Getenv("PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := cache.NewCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(c.Set(ctx, &cache.Item{Key: "foo", Value: "dfgjkng"}))
	fmt.Println(c.Set(ctx, &cache.Item{Key: "asd", Value: "12312312"}))

	fmt.Println("-----")

	fmt.Println(c.Get(ctx, &cache.Key{Key: "foo"}))
	fmt.Println(c.Get(ctx, &cache.Key{Key: "asd"}))
}
