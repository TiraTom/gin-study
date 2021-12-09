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

func (c *CreateTask) Do(p *gr.CreateTaskRequestParam) (*gr.Task, error) {
	newTask, err := domain_obj.NewTaskToCreate(p)
	if err != nil {
		return nil, fmt.Errorf("タスク作成のパラメーター処理においてエラーが発生しました(p={Name=%v Details=%v ImportanceName=%v Deadline=%v}): %w", p.Name, p.Details, p.ImportanceName, p.Deadline, err)
	}

	createdTask, err := c.tr.Create(newTask)
	if err != nil {
		return nil, fmt.Errorf("タスク作成においてエラーが発生しました(p={Name=%v Details=%v ImportanceName=%v Deadline=%v}): %w", newTask.Name, newTask.Details, newTask.ImportanceName, newTask.Deadline, err)
	}

	t, err := createdTask.ToDto()
	if err != nil {
		return nil, fmt.Errorf("タスク作成成功後戻り値生成においてエラーが発生しました(p={Name=%v Details=%v ImportanceName=%v Deadline=%v}): %w", createdTask.Name, createdTask.Details, createdTask.ImportanceName, createdTask.Deadline, err)
	}

	return t, err
}

func NewCreateTask(tr repository_interface.Task) *CreateTask {
	return &CreateTask{tr}
}
