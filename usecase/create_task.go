package usecase

import (
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type CreateTask struct {
	tr repository_interface.Task
}

func (c *CreateTask) Do(p *gr.CreateTaskRequestParam) (*gr.Task, error) {
	newTask, err := domain_obj.NewTaskToCreate(p)
	if err != nil {
		return nil, err
	}

	createdTask, err := c.tr.Create(newTask)
	if err != nil {
		return nil, err
	}

	return createdTask.ToDto()
}

func NewCreateTask(tr repository_interface.Task) *CreateTask {
	return &CreateTask{tr}
}
