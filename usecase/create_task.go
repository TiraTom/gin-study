package usecase

import (
	"fmt"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type CreateTask struct {
	tr repository_interface.Task
}

func (c *CreateTask) Do(param *gr.CreateTaskRequestParam) (*gr.Task, error) {
	// TODO バリデーション

	newTask, err := domain_obj.NewTaskToCreate(param)
	if err != nil {
		return nil, err
	}

	err = c.tr.Create(newTask)
	if err != nil {
		return nil, err
	}

	createdTask, err := newTask.ToDto()
	if err != nil {
		return nil, fmt.Errorf("データ登録成功後、内部エラーが発生しました %w", err)
	}

	return createdTask, nil
}

func NewCreateTask(tr repository_interface.Task) *CreateTask {
	return &CreateTask{tr}
}
