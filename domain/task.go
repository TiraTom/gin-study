package domain

import (
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Task struct {
	Id           string
	Name         string
	Details      string
	Importance   Importance
	RegisteredAt time.Time `gorm:"createdAt"`
	Deadline     time.Time
	UpdatedAt    time.Time
}

func (t *Task) ToDto() (*gr.Task, error) {
	importance, err := t.Importance.ToDto()
	if err != nil {
		return nil, err
	}

	return &gr.Task{
		Id:           t.Id,
		Name:         t.Name,
		Details:      t.Details,
		Importance:   importance,
		RegisteredAt: timestamppb.New(t.RegisteredAt),
		Deadline:     timestamppb.New(t.Deadline),
		UpdatedAt:    timestamppb.New(t.UpdatedAt),
	}, nil
}

func NewTask(tr *infrastructure.TaskAndImportanceRecord) *Task {
	return &Task{
		Id:      tr.Id,
		Name:    tr.Name,
		Details: tr.Details,
		Importance: Importance{
			Name:  tr.ImportanceName,
			Level: tr.ImportanceLevel,
		},
		RegisteredAt: tr.RegisteredAt,
		Deadline:     tr.Deadline,
		UpdatedAt:    tr.UpdatedAt,
	}
}
