package usecase

import (
	"github.com/Tiratom/gin-study/domain/repository_interface"

	gr "github.com/Tiratom/gin-study/grpc"
)

type DeleteTask struct {
	tr repository_interface.Task
}

func (d *DeleteTask) Do(p *gr.DeleteTaskRequestParam) error {
	return d.tr.Delete(p.Id)
}

func NewDeleteTask(tr repository_interface.Task) *DeleteTask {
	return &DeleteTask{tr}
}
