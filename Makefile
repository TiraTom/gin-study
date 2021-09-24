# .protoファイルをコンパイル
grpc: FORCE
	protoc --go_out=plugins=grpc:./gRPC --go_opt=module=github.com/Tiratom/gin-study/grpc ./gRPC/gin-study.proto

grpcDoc: FORCE
	protoc --doc_out=./doc --doc_opt=markdown,api.md ./gRPC/gin-study.proto

# ローカル環境設定でのサーバー起動
run:
	ENV=local go run main.go

grpc-list:
	grpcurl -plaintext localhost:8081 list

# gRPCサーバーにお試しリクエストを送る（main.go内でリフレクションサービスを有効にしていない版）
hello1:
	grpcurl -plaintext -import-path . -proto ./gRPC/gin-study.proto localhost:8081 TaskService/GetAllTasks
# gRPCサーバーにお試しリクエストを送る（main.go内でリフレクションサービスを有効にしている版）
getAllTasks:
	grpcurl -plaintext localhost:8081 TaskService/GetAllTasks
# gRPCサーバーにお試しリクエストを送る
createATask:
	grpcurl -plaintext -d '{"name": "TestTask1", "details": "TestDetails1", "importanceName": "LOW", "deadline": "2021-09-23T14:30:00+09:00"}' localhost:8081 TaskService/CreateTask

# DI用ファイル作成
di: FORCE
	wire di/wire.go

lint:
	golangci-lint run

# ローカル開発用mysqlの立ち上げ
local-db:
	docker compose up

FORCE: