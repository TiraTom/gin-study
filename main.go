package main

import (
	"net"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Ginとしては動かさない（gRPC利用）のでコメントアウト
	// r := config.GetRouter()
	// r.Run(":8080")

	// middlewareの設定
	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_zap.StreamServerInterceptor(middleware.GetZapLogger()),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(middleware.GetZapLogger()),
			),
		),
	)

	// サービスの登録
	gr.RegisterCatServiceServer(server, &presentation.CatServiceServer{})
	gr.RegisterTaskServiceServer(server, &presentation.TaskServiceServer{})

	// grpcurlコマンドで叩けるようにリフレクションサービスを登録
	reflection.Register(server)

	// gRPCサーバーの起動設定
	listenPort, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	err = server.Serve(listenPort)
	if err != nil {
		panic(err)
	}
}
