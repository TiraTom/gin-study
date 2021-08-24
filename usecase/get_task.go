package usecase

import (
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type GetTask struct {
	tr repository_interface.Task
}

func (gt *GetTask) GetAllTasks() (*gr.Tasks, error) {
	allTasks := gt.tr.GetAll()
	return (&domain_obj.Tasks{Value: allTasks}).ToDto()
}

func NewGetTask(tr repository_interface.Task) *GetTask {
	return &GetTask{tr}
}
