package main

import (
	"log"
	"net"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/presentation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Ginとしては動かさない（gRPC利用）のでコメントアウト
	// r := config.GetRouter()
	// r.Run(":8080")

	listenPort, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	gr.RegisterCatServiceServer(server, &presentation.CatServiceServer{})
	gr.RegisterTaskServiceServer(server, &presentation.TaskServiceServer{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	err = server.Serve(listenPort)
	if err != nil {
		log.Fatal(err)
	}
}
