//+build wireinject

package di

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	"github.com/Tiratom/gin-study/usecase/usecase_interface"
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

func InitializeImportanceRepositoryInterface() repository_interface.Importance {
	wire.Build(repository_interface.NewImportance, InitializeDB)
	return nil
}

func InitializeTaskRepositoryInterface() repository_interface.Task {
	wire.Build(repository_interface.NewTask, InitializeDB)
	return nil
}

func InitializeGetTaskUsecaseIF() usecase_interface.GetTask {
	wire.Build(usecase_interface.NewGetTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeCreateTaskUsecaseIF() usecase_interface.CreateTask {
	wire.Build(usecase_interface.NewCreateTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeUpdateTaskUsecaseIF() usecase_interface.UpdateTask {
	wire.Build(usecase_interface.NewUpdateTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeDeleteTaskUsecaseIF() usecase_interface.DeleteTask {
	wire.Build(usecase_interface.NewDeleteTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeSearchTaskUsecaseIF() usecase_interface.SearchTask {
	wire.Build(usecase_interface.NewSearchTask, InitializeTaskRepositoryInterface)
	return nil
}

func InitializeTaskServiceServer() *presentation.TaskServiceServer {
	wire.Build(presentation.NewTaskServiceServer, middleware.NewZapLogger, InitializeGetTaskUsecaseIF, InitializeCreateTaskUsecaseIF, InitializeUpdateTaskUsecaseIF, InitializeDeleteTaskUsecaseIF, InitializeSearchTaskUsecaseIF)
	return nil
}
