package domain_obj

import (
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/infrastructure/record"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Task struct {
	Id             string
	Name           string
	Details        string
	ImportanceName string
	RegisteredAt   time.Time `gorm:"autoCreateTime"`
	Deadline       time.Time
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	Version        uint
}

func (t *Task) ToDto() (*gr.Task, error) {
	return &gr.Task{
		Id:             t.Id,
		Name:           t.Name,
		Details:        t.Details,
		ImportanceName: t.ImportanceName,
		RegisteredAt:   timestamppb.New(t.RegisteredAt),
		Deadline:       timestamppb.New(t.Deadline),
		UpdatedAt:      timestamppb.New(t.UpdatedAt),
	}, nil
}

func (t *Task) ToRecord(i int64) *record.Task {
	return &record.Task{
		Id:           t.Id,
		Name:         t.Name,
		Details:      t.Details,
		ImportanceId: i,
		RegisteredAt: t.RegisteredAt,
		Deadline:     t.Deadline,
		UpdatedAt:    t.UpdatedAt,
		Version:      t.Version,
	}
}

func NewTask(tr *record.TaskAndImportance) *Task {
	return &Task{
		Id:             tr.Id,
		Name:           tr.Name,
		Details:        tr.Details,
		ImportanceName: tr.ImportanceName,
		RegisteredAt:   tr.RegisteredAt,
		Deadline:       tr.Deadline,
		UpdatedAt:      tr.UpdatedAt,
	}
}

// NewTaskToCreate　リクエストパラムから新規作成するタスクの定義を用意する。
func NewTaskToCreate(p *gr.CreateTaskRequestParam) (*Task, error) {
	now := time.Now()

	return &Task{
		Id:             uuid.New().String(),
		Name:           p.Name,
		Details:        p.Details,
		ImportanceName: p.ImportanceName,
		RegisteredAt:   now,
		Deadline:       p.Deadline.AsTime(),
		UpdatedAt:      now,
		Version:        1,
	}, nil
}