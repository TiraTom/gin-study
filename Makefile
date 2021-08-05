# .protoファイルをコンパイル
grpc: 
	protoc --go_out=plugins=grpc:./gRPC ./gRPC/gin-study.proto