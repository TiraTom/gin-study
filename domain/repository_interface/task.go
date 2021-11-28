package repository_interface

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/infrastructure/repository_impl"
)

type Task interface {
	GetAll() (*domain_obj.Tasks, error)
	GetById(id string) (*domain_obj.Task, error)
	Create(*domain_obj.Task) (*domain_obj.Task, error)
	Update(*domain_obj.Task) (*domain_obj.Task, error)
	Delete(id string) error
	Search(*domain_obj.TaskSearchCondition) (*domain_obj.Tasks, error)
}

func NewTask(db *config.DB) Task {
	return repository_impl.NewTask(db)
}
