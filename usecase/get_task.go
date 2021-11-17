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
	allTasks := gt.tr.GetAll()

	// TODO 例外処理の実装が終わったら削除する
	panic(fmt.Errorf("recoveryのテスト"))

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
