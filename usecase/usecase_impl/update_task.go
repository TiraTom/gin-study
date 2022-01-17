package usecase_impl

import (
	"fmt"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type UpdateTask struct {
	Tr repository_interface.Task
}

func (u *UpdateTask) Do(p *gr.UpdateTaskRequestParam) (*domain_obj.Task, error) {
	// memo: RESTapiだとidと更新内容のパラメーターは別変数でもらうだろうから、この引数の書き方だとpresentation層の実装の影響をusecaseが受けてしまうような気もする、、
  // →このメソッドの引数はドメインオブジェクトにして、controller側でリクエストパラメーターをドメインオブジェクトに詰め替えるというようにした方が分離ができそうだ

	targetTask, err := u.Tr.GetById(p.Id)
	if err != nil {
		return nil, fmt.Errorf("更新対象タスク取得においてエラーが発生しました(id=%v); %w", p.Id, err)
	}

	if !targetTask.IsNeededToUpdate(p) {
		// 更新項目なしなので早期リターン
		return targetTask, nil
	}

	newTaskToUpdate, err := domain_obj.NewTaskToUpdate(targetTask, p)
	if err != nil {
		return nil, fmt.Errorf("更新対象タスクのパラメーター処理においてエラーが発生しました({Id=%v Name=%v Details=%v ImportanceName=%v Deadline=%v}); %w", p.Id, p.Name, p.Details, p.ImportanceName, p.Deadline, err)
	}

	updateResult, err := u.Tr.Update(newTaskToUpdate)
	if err != nil {
		return nil, fmt.Errorf("タスク更新においてエラーが発生しました({Id=%v Name=%v Details=%v ImportanceName=%v Deadline=%v}); %w", p.Id, p.Name, p.Details, p.ImportanceName, p.Deadline, err)
	}

	return updateResult, err
}

func NewUpdateTask(tr repository_interface.Task) *UpdateTask {
	return &UpdateTask{tr}
}
