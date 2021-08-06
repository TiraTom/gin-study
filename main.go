package main

import (
	"log"
	"net"

	"github.com/Tiratom/gin-study/controller"
	gr "github.com/Tiratom/gin-study/grpc"
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
	gr.RegisterCatServer(server, &controller.CatServer{})
	gr.RegisterTaskServiceServer(server, &controller.TaskServiceServer{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	err = server.Serve(listenPort)
	if err != nil {
		log.Fatal(err)
	}
}
