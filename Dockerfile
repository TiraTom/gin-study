FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /build-output

# FROM gcr.io/distroless/base:debug as runner → ELF: not foundとかsyntax error: unexpected end of file (expecting ")")
FROM gcr.io/distroless/base as runner

COPY --from=builder /build-output /app/gin-study
COPY --from=builder /go/src/conf-files /app/conf-files
COPY --from=builder /go/src/migrations /app/migrations

WORKDIR /app
CMD ["./gin-study"]