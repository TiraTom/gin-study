package usecase

import (
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type UpdateTask struct {
	tr repository_interface.Task
}

func (u *UpdateTask) Do(p *gr.UpdateTaskRequestParam) (*domain_obj.Task, error) {
	// memo: RESTapiだとidと更新内容のパラメーターは別変数でもらうだろうから、この引数の書き方だとpresentation層の実装の影響をusecaseが受けてしまうような気もする、、

	targetTask, err := u.tr.GetById(p.Id)
	if err != nil {
		return nil, err
	}

	newTaskToUpdate, err := domain_obj.NewTaskToUpdate(targetTask, p)
	if err != nil {
		return nil, err
	}

	return u.tr.Update(newTaskToUpdate)
}

func NewUpdateTask(tr repository_interface.Task) *UpdateTask {
	return &UpdateTask{tr}
}
