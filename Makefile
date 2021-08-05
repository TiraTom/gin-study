# .protoファイルをコンパイル
grpc: 
	protoc --go_out=plugins=grpc:./gRPC --go_opt=module=github.com/Tiratom/gin-study/grpc ./gRPC/gin-study.proto

# サーバー起動
run:
	go run main.go