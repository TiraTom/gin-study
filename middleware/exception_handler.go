package middleware

import (
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandlePanic(p interface{}) error {
	zap.L().Error(fmt.Sprintf("%+v\n", p))
	return status.Errorf(codes.Internal, "Unexpected error")
}
