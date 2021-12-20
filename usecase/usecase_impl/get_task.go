package usecase_impl

import (
	"fmt"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
)

type GetTask struct {
	Tr repository_interface.Task
}

func (gt *GetTask) DoAll() (*domain_obj.Tasks, error) {
	allTasks, err := gt.Tr.GetAll()
	if err != nil {
		return nil, fmt.Errorf("タスク全件取得においてエラーが発生しました; %w", err)
	}
	return allTasks, err
}

func (gt *GetTask) DoById(id string) (*domain_obj.Task, error) {
	task, err := gt.Tr.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("idによるタスク取得においてエラーが発生しました(id=%v); %w", id, err)
	}
	return task, err
}

func NewGetTask(tr repository_interface.Task) *GetTask {
	return &GetTask{tr}
}
