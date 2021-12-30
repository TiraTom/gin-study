FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# golangのイメージでそのまま動かすパターン↓
# RUN go build -o /app/gin-study
# COPY conf-files /app/
# COPY migrations /app/
# CMD ["/app/gin-study"]
# golangのイメージでそのまま動かすパターン↑

RUN go build -o /build-output

# FROM gcr.io/distroless/base-debian11:debug as runner
FROM ubuntu:latest as runner

# COPY --from=builder /go/src/conf-files /go/src/migrations ./
# COPY --from=builder /go/src/conf-files /app/
# COPY --from=builder /go/src/conf-files/zap_config.json /app/conf-files/zap_config.json→OK
# COPY --from=builder /go/src/migrations /app/→フォルダ内のファイルが展開される

# COPY --from=builder /build-output /app/gin-study
# COPY --from=builder /go/src/conf-files /app/conf-files
# COPY --from=builder /go/src/migrations /app/migrations

COPY --from=builder /build-output /gin-study
COPY --from=builder /go/src/conf-files /conf-files
COPY --from=builder /go/src/migrations /migrations

# ポートはdocker runコマンドで指定しようかな、、その時に応じて変わるだろうし
EXPOSE 8081
# EXPOSE 3306

# WORKDIR /app
# CMD ["/app"]
# CMD ["./gin-study"]
CMD ["/gin-study"]
