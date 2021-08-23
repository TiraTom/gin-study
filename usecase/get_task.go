package usecase

import (
	"github.com/Tiratom/gin-study/domain"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/repository_interface"
)

type GetTask struct {
	tr repository_interface.Task
}

func (gt *GetTask) GetAllTasks() (*gr.Tasks, error) {
	allTasks := gt.tr.GetAll()
	return (&domain.Tasks{Value: allTasks}).ToDto()
}

func NewGetTask(tr repository_interface.Task) *GetTask {
	return &GetTask{tr}
}
