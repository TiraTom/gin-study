package usecase_interface

import (
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/usecase/usecase_impl"
)

type CreateTask interface {
	Do(p *gr.CreateTaskRequestParam) (*domain_obj.Task, error)
}

func NewCreateTask(tr repository_interface.Task) CreateTask {
	return usecase_impl.NewCreateTask(tr)
}

type DeleteTask interface {
	Do(p *gr.DeleteTaskRequestParam) error
}

func NewDeleteTask(tr repository_interface.Task) DeleteTask {
	return usecase_impl.NewDeleteTask(tr)
}

type GetTask interface {
	DoAll() (*domain_obj.Tasks, error)
	DoById(id string) (*domain_obj.Task, error)
}

func NewGetTask(tr repository_interface.Task) GetTask {
	return usecase_impl.NewGetTask(tr)
}

type SearchTask interface {
	Do(p *gr.GetTaskByConditionRequestParam) (*domain_obj.Tasks, error)
}

func NewSearchTask(tr repository_interface.Task) SearchTask {
	return usecase_impl.NewSearchTask(tr)
}

type UpdateTask interface {
	Do(p *gr.UpdateTaskRequestParam) (*domain_obj.Task, error)
}

func NewUpdateTask(tr repository_interface.Task) UpdateTask {
	return usecase_impl.NewUpdateTask(tr)
}
