package main

import (
	"fmt"
	"net"

	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/di"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Ginとしては動かさない（gRPC利用）のでコメントアウト
	// r := config.GetRouter()
	// r.Run(":8080")

	// グローバルロガーの設定
	zap.ReplaceGlobals(middleware.GetZapLogger())

	// middlewareの設定
	server := grpc.NewServer(
		// TODO: 処理終わりにリクエストIDなしのログが出てしまうので一旦コメントアウト
		// grpc.StreamInterceptor(
		// 	grpc_middleware.ChainStreamServer(
		// 		grpc_zap.StreamServerInterceptor(),
		// 	),
		// ),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.GetZapLoggerUnaryInterceptor(),
			),
		),
	)

	// サービスの登録
	gr.RegisterCatServiceServer(server, &presentation.CatServiceServer{})
	gr.RegisterTaskServiceServer(server, di.InitializeTaskServiceServer())

	// grpcurlコマンドで叩けるようにリフレクションサービスを登録
	reflection.Register(server)

	// 環境変数読み込み
	env := config.NewEnvironment()

	// gRPCサーバーの起動設定
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%s", env.APP_PORT_NUM))
	if err != nil {
		panic(err)
	}

	err = server.Serve(listenPort)
	if err != nil {
		panic(err)
	}
}
