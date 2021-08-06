package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	gr "github.com/Tiratom/gin-study/grpc"
)

// main grpcurlコマンドを使わずにリクエストを送りたいとき用のClientとしてのmainメソッド
func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error")
	}
	defer conn.Close()

	client := gr.NewCatClient(conn)
	message := &gr.GetMyCatMessage{TargetCat: "tama"}
	res, err := client.GetMyCat(context.TODO(), message)

	if res != nil {
		fmt.Printf("result:%#v \n", res)
	}
	if err != nil {
		fmt.Printf("error::%#v \n", err.Error())
	}
}
