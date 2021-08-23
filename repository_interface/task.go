package repository_interface

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/repository"
)

type Task interface {
	GetAll() []*domain.Task
}

func NewTask(db *config.DB) Task {
	return infrastructure.NewTaskRepository(db)
}
