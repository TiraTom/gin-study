package domain

import (
	gr "github.com/Tiratom/gin-study/grpc"
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
