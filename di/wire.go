//+build wireinject

package di

import (
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	"github.com/google/wire"
)

func InitializeTaskServiceServer() *presentation.TaskServiceServer {
	wire.Build(presentation.NewTaskServiceServer, middleware.NewZapLogger)
	return nil
}
