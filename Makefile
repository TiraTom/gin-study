# .protoファイルをコンパイル
grpc: FORCE
	protoc --go_out=plugins=grpc:./gRPC --go_opt=module=github.com/Tiratom/gin-study/grpc ./gRPC/gin-study.proto

grpcDoc: FORCE
	protoc --doc_out=./doc --doc_opt=markdown,api.md ./gRPC/gin-study.proto

# サーバー起動
run:
	go run main.go

# gRPCサーバーにお試しリクエストを送る（main.go内でリフレクションサービスを有効にしていない版）
hello1:
	grpcurl -plaintext -import-path . -proto ./gRPC/gin-study.proto localhost:8081 TaskService/GetAllTasks
# gRPCサーバーにお試しリクエストを送る（main.go内でリフレクションサービスを有効にしていない版）
hello2:
	grpcurl -plaintext localhost:8081 TaskService/GetAllTasks

# DI用ファイル作成
di: FORCE
	wire di/wire.go

FORCE: