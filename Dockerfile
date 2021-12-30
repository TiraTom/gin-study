FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /build-output

FROM ubuntu:latest as runner
# FROM gcr.io/distroless/base-debian11:debug as runner

WORKDIR /app
COPY --from=builder /build-output /app/gin-study
COPY --from=builder /go/src/conf-files /app/conf-files
COPY --from=builder /go/src/migrations /app/migrations

CMD ["./gin-study"]
