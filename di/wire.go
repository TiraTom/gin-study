//+build wireinject

package di

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	"github.com/Tiratom/gin-study/infrastructure/repository_impl"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	"github.com/Tiratom/gin-study/usecase"
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

func InitializeImportanceRepository() *repository_impl.Importance {
	wire.Build(repository_impl.NewImportance, InitializeDB)
	return nil
}

func InitializeImportanceRepositoryInterface() repository_interface.Importance {
	wire.Build(repository_interface.NewImportance, InitializeDB)
	return nil
}

func InitializeTaskRepositoryInterface() repository_interface.Task {
	wire.Build(repository_interface.NewTask, InitializeDB)
	return nil
}

func InitializeGetTaskUsecase() *usecase.GetTask {
	wire.Build(usecase.NewGetTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeCreateTaskUsecase() *usecase.CreateTask {
	wire.Build(usecase.NewCreateTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeUpdateTaskUsecase() *usecase.UpdateTask {
	wire.Build(usecase.NewUpdateTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeTaskServiceServer() *presentation.TaskServiceServer {
	wire.Build(presentation.NewTaskServiceServer, middleware.NewZapLogger, InitializeGetTaskUsecase, InitializeCreateTaskUsecase, InitializeUpdateTaskUsecase)
	return nil
}
