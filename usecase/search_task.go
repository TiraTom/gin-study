package usecase

import (
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
)

type SearchTask struct {
	tr repository_interface.Task
}

func (s *SearchTask) Do(p *gr.GetTaskByConditionRequestParam) (*gr.Tasks, error) {
	tasks, err := s.tr.Search(p)
	if err != nil {
		return nil, err
	}
	return (&domain_obj.Tasks{Value: tasks}).ToDto()
}

func NewSearchTask(tr repository_interface.Task) *SearchTask {
	return &SearchTask{tr}
}
