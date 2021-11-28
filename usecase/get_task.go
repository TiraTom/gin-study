package usecase

import (
	"fmt"

	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type GetTask struct {
	tr repository_interface.Task
}

func (gt *GetTask) DoAll() (*gr.Tasks, error) {
	allTasks, err := gt.tr.GetAll()
	if err != nil {
		return nil, fmt.Errorf("タスク取得においてエラーが発生しました %w", err)
	}
	return allTasks.ToDto()
}

func (gt *GetTask) DoById(id string) (*gr.Task, error) {
	task, err := gt.tr.GetById(id)
	if err != nil {
		return nil, err
	}
	return task.ToDto()
}

func NewGetTask(tr repository_interface.Task) *GetTask {
	return &GetTask{tr}
}
