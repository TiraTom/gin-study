package controller

import (
	"context"

	pb "github.com/Tiratom/gin-study/grpc"
)

type CatServer struct {
}

func (c *CatServer) GetMyCat(ctx context.Context, msg *pb.GetMyCatMessage) (*pb.MyCatResponse, error)
