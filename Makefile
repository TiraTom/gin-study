# .protoファイルをコンパイル
grpc: 
	protoc --go_out=./controller/gRPC --go_opt=paths=source_relative --go-grpc_out=./controller/gRPC --go-grpc_opt=paths=source_relative ./controller/gRPC/gin-study.proto