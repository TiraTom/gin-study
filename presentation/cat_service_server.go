package presentation

import (
	"context"
	"errors"

	pb "github.com/Tiratom/gin-study/grpc"
)

type CatServiceServer struct {
}

// https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3　参考にお試し実装
func (c *CatServiceServer) GetMyCat(ctx context.Context, msg *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	switch msg.TargetCat {
	case "tama":
		//たまはメインクーン
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		//ミケはノルウェージャンフォレストキャット
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("not found your cat")
}
