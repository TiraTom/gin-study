package usecase

import (
	"fmt"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type SearchTask struct {
	tr repository_interface.Task
}

func (s *SearchTask) Do(p *gr.GetTaskByConditionRequestParam) (*gr.Tasks, error) {
	c := domain_obj.NewTaskSearchCondition(p)
	tasks, err := s.tr.Search(c)
	if err != nil {
		return nil, fmt.Errorf("タスク検索の検索条件処理においてエラーが発生しました(検索条件={Name=%v Details=%v ImportanceName=%v Deadline=%v SearchTypeForDeadline=%v}): %w", p.Name, p.Details, p.ImportanceName, p.Deadline, p.SearchTypeForDeadline, err)
	}

	t, err := tasks.ToDto()
	if err != nil {
		return nil, fmt.Errorf("タスク検索結果の変換処理においてエラーが発生しました(検索結果=%v):%w", t, err)
	}

	return t, err
}

func NewSearchTask(tr repository_interface.Task) *SearchTask {
	return &SearchTask{tr}
}
