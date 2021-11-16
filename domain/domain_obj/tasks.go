package domain_obj

import (
	gr "github.com/Tiratom/gin-study/grpc"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
)

type Tasks struct {
	Value []*Task
}

func (t *Tasks) ToDto() (*gr.Tasks, error) {
	taskDtos := make([]*gr.Task, len(t.Value))

	for i, v := range t.Value {
		dto, err := v.ToDto()
		if err != nil {
			return nil, err
		}

		taskDtos[i] = dto
	}

	return &gr.Tasks{Tasks: taskDtos}, nil
}

func NewTasks(records []*infrastructure.TaskAndImportance) *Tasks {
	values := make([]*Task, len(records))
	for i, v := range records {
		values[i] = NewTask(v)
	}

	return &Tasks{values}
}
