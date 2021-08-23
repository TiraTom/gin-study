//+build wireinject

package di

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/infrastructure"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	"github.com/Tiratom/gin-study/repository_interface"
	"github.com/google/wire"
)

func InitializeEnvironment() *config.Environment {
	wire.Build(config.NewEnvironment)
	return nil
}

func InitializeDB() *config.DB {
	wire.Build(config.NewDB, config.NewEnvironment)
	return nil
}

func InitializeImportanceRepository() *infrastructure.ImportanceRepository {
	wire.Build(infrastructure.NewImportanceRepository, InitializeDB)
	return nil
}

func InitializeImportanceRepositoryInterface() repository_interface.Importance {
	wire.Build(repository_interface.NewImportance, InitializeDB)
	return nil
}

func InitializeTaskServiceServer() *presentation.TaskServiceServer {
	wire.Build(presentation.NewTaskServiceServer, middleware.NewZapLogger, InitializeImportanceRepositoryInterface)
	return nil
}
