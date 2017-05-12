package main

import (
	"log"
	"time"

	pb "./pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewServerClient(conn)
	stream, err := client.Start(context.Background())
	if err != nil {
		panic(err)
	}

	resp, err := stream.Recv()
	if err != nil {
		panic(err)
	}
	log.Println("recv resp:", resp.Id)

	time.Sleep(time.Second * 5)
	log.Println("sleep done")

	conn.Close()
}
