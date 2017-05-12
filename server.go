package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	. "./pb"
)

type server struct{}

func (s *server) Start(stream Server_StartServer) error {
	log.Println("start")
	stream.Send(&StartResp{
		Id: "asdf",
	})
	log.Println("send resp done")

	go func() {
		log.Println("before recv")
		_, err := stream.Recv()
		log.Println("recv done")
		log.Println(err)
	}()

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterServerServer(grpcServer, new(server))

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(listen)
}
