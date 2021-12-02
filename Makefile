# .protoファイルをコンパイル
grpc: FORCE
	protoc --proto_path=${GOPATH}/src --proto_path=:./grpc --go_out=plugins=grpc:./grpc --go_opt=module=github.com/Tiratom/gin-study/grpc --govalidators_out=paths=source_relative:./grpc ./grpc/gin-study.proto

grpcDoc: FORCE
	protoc --doc_out=./doc --doc_opt=markdown,api.md ./gRPC/gin-study.proto

# ローカル環境設定でのサーバー起動
run:
	ENV=local go run main.go

test:
	ENV=test go test -v ./...

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
	grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext -d '{"id": "1dd2a403-5193-41b4-8462-ab9a4a9385e6", "importanceName": "HIGH", "deadline": "2021-10-23T14:30:00+09:00"}' localhost:8081 TaskService/UpdateTask
# gRPCサーバーにタスク削除のお試しリクエストを送る
deleteATask:
	grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext -d '{"id": "2"}' localhost:8081 TaskService/DeleteTask
# gRPCサーバーにタスク取得のお試しリクエストを送る
getTask:
	grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext -d '{"id": "1dd2a403-5193-41b4-8462-ab9a4a9385e6"}' localhost:8081 TaskService/GetTask
# gRPCサーバーにタスク検索のお試しリクエストを送る
getTasks:
	grpcurl -import-path . -proto ./grpc/gin-study.proto -import-path ${GOPATH}/src -proto github.com/mwitkow/go-proto-validators/validator.proto -plaintext -d '{"importanceName":"LOW","deadline": "2021-08-20T09:25:53Z", "searchTypeForDeadline": "BEFORE"}' localhost:8081 TaskService/GetTasks

# DI用ファイル作成
di: FORCE
	wire di/wire.go

lint:
	golangci-lint run

# ローカル開発用mysqlの立ち上げ（テスト用DB含む）
local-db:
	docker compose up

insert-dummyData:
	QUERY="$$(cat ./dummyData/dummyData.sql | tr -d '\n')"; \
	docker exec -it db mysql -uroot -p -h 127.0.0.1 -p -Dgin_study -e "$$QUERY" -p

show-migrate-ver:
	migrate -source file://migrations -database "mysql://docker:docker@tcp(localhost:3306)/gin_study" version

# migrate時にエラーによりdirtyになった場合に、マイグレーションファイル修正後に指定VerでDirty状態を解除するためのコマンドメモ（VERSIONを適宜書き換えること）
migrate-force:
	migrate -path ./migrations -database "mysql://docker:docker@tcp(localhost:3306)/gin_study" force 【VERSION】

# schemaspyによるDBスキーマの作成
schema:
	docker run --rm --network=host -v "$$(PWD)/schemaspy/output:/output" -v "$$(PWD)/schemaspy/schemaspy.properties:/schemaspy.properties" -v "$$(PWD)/schemaspy/drivers:/drivers" schemaspy/schemaspy:latest -debug -connprops "useSSL\=false;allowPublicKeyRetrieval\=true"

FORCE: