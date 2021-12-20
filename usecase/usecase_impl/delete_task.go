package usecase_impl

import (
	"fmt"

	"github.com/Tiratom/gin-study/domain/repository_interface"

	gr "github.com/Tiratom/gin-study/grpc"
)

type DeleteTask struct {
	Tr repository_interface.Task
}

func (d *DeleteTask) Do(p *gr.DeleteTaskRequestParam) error {
	err := d.Tr.Delete(p.Id)
	if err != nil {
		return fmt.Errorf("タスク削除においてエラーが発生しました(id=%v); %w", p.Id, err)
	}
	return err
}

func NewDeleteTask(tr repository_interface.Task) *DeleteTask {
	return &DeleteTask{tr}
}
