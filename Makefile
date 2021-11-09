# .protoファイルをコンパイル
grpc: FORCE
	protoc --proto_path=${GOPATH}/src --proto_path=:./grpc --go_out=plugins=grpc:./grpc --go_opt=module=github.com/Tiratom/gin-study/grpc --govalidators_out=paths=source_relative:./grpc ./grpc/gin-study.proto

grpcDoc: FORCE
	protoc --doc_out=./doc --doc_opt=markdown,api.md ./gRPC/gin-study.proto

# ローカル環境設定でのサーバー起動
run:
	ENV=local go run main.go

grpc-list:
	grpcurl -plaintext localhost:8081 list

# gRPCサーバーにお試しリクエストを送る（main.go内でリフレクションサービスを有効にしていない版）
hello1:
	grpcurl -plaintext -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto localhost:8081 TaskService/GetAllTasks
# gRPCサーバーにお試しリクエストを送る（main.go内でリフレクションサービスを有効にしている版）
# ※go-proto-validator利用時はリフレクションサービスが使えないので-import-path, -protoを指定している。参考：<https://qiita.com/gold-kou/items/4e17f98976b43433fa8d>
getAllTasks:
	grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext localhost:8081 TaskService/GetAllTasks
# gRPCサーバーにタスク作成のお試しリクエストを送る
createATask:
	grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext -d '{"name": "TestTask1", "details": "TestDetails1", "importanceName": "LOW", "deadline": "2021-09-23T14:30:00+09:00"}' localhost:8081 TaskService/CreateTask
# gRPCサーバーにタスク更新のお試しリクエストを送る
updateATask:
	grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext -d '{"id": "1", "importanceName": "HIGH"}' localhost:8081 TaskService/UpdateTask
# grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext -d '{"id": "1", "name": "TestTask1ver3", "deadline": "2021-11-10T00:00:00+09:00"}' localhost:8081 TaskService/UpdateTask

# DI用ファイル作成
di: FORCE
	wire di/wire.go

lint:
	golangci-lint run

# ローカル開発用mysqlの立ち上げ
local-db:
	docker compose up

FORCE: