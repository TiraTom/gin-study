package domain_obj

import (
	"fmt"
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
	Deadline       *time.Time
	RegisteredAt   *time.Time
	UpdatedAt      *time.Time
	Version        uint
}

var nowTimeFunc = time.Now
var newUuidFunc = uuid.New

func (t *Task) ToDto() (*gr.Task, error) {
	var ra *timestamppb.Timestamp
	if t.RegisteredAt != nil {
		ra = timestamppb.New(*t.RegisteredAt)
	} else {
		ra = nil
	}

	var dl *timestamppb.Timestamp
	if t.Deadline != nil {
		dl = timestamppb.New(*t.Deadline)
	} else {
		dl = nil
	}

	var ua *timestamppb.Timestamp
	if t.UpdatedAt != nil {
		ua = timestamppb.New(*t.UpdatedAt)
	} else {
		ua = nil
	}

	return &gr.Task{
		Id:             t.Id,
		Name:           t.Name,
		Details:        t.Details,
		ImportanceName: t.ImportanceName,
		RegisteredAt:   ra,
		Deadline:       dl,
		UpdatedAt:      ua,
	}, nil
}

func (t *Task) ToRecord(i int64) *record.Task {
	var ra time.Time
	if t.RegisteredAt == nil {
		ra = *new(time.Time)
	} else {
		ra = *t.RegisteredAt
	}

	var dl time.Time
	if t.Deadline == nil {
		dl = *new(time.Time)
	} else {
		dl = *t.Deadline
	}

	var ua time.Time
	if t.UpdatedAt == nil {
		ua = *new(time.Time)
	} else {
		ua = *t.UpdatedAt
	}

	return &record.Task{
		Id:             t.Id,
		Name:           t.Name,
		Details:        t.Details,
		ImportanceId:   i,
		RegisteredTime: ra,
		Deadline:       dl,
		UpdatedTime:    ua,
		Version:        t.Version,
	}
}

func NewTask(tr *record.TaskAndImportance) *Task {
	return &Task{
		Id:             tr.Id,
		Name:           tr.Name,
		Details:        tr.Details,
		ImportanceName: tr.ImportanceName,
		RegisteredAt:   &tr.RegisteredTime,
		Deadline:       &tr.Deadline,
		UpdatedAt:      &tr.UpdatedTime,
		Version:        tr.Version,
	}
}

// NewTaskToCreate　リクエストパラムから新規作成するタスクの定義を用意する。
func NewTaskToCreate(p *gr.CreateTaskRequestParam) (*Task, error) {
	now := nowTimeFunc().UTC()

	newDeadline := p.Deadline.AsTime()

	return &Task{
		Id:             newUuidFunc().String(),
		Name:           p.Name,
		Details:        p.Details,
		ImportanceName: p.ImportanceName,
		RegisteredAt:   &now,
		Deadline:       &newDeadline,
		UpdatedAt:      &now,
		Version:        1,
	}, nil
}

// NewTaskToUpdate　リクエストパラムから更新保存したいタスクの定義を用意する。
func NewTaskToUpdate(o *Task, p *gr.UpdateTaskRequestParam) (*Task, error) {
	now := nowTimeFunc().UTC()

	var newName string
	if p.Name != "" {
		newName = p.Name
	} else {
		newName = o.Name
	}

	var newDetails string
	if p.Details != "" {
		newDetails = p.Details
	} else {
		newDetails = o.Details
	}

	var newImportanceName string
	if p.ImportanceName != "" {
		newImportanceName = p.ImportanceName
	} else {
		newImportanceName = o.ImportanceName
	}

	var newDeadline time.Time
	if p.Deadline != nil {
		newDeadline = p.Deadline.AsTime()
	} else {
		newDeadline = *o.Deadline
	}

	return &Task{
		Id:             o.Id,
		Name:           newName,
		Details:        newDetails,
		ImportanceName: newImportanceName,
		RegisteredAt:   o.RegisteredAt,
		Deadline:       &newDeadline,
		UpdatedAt:      &now,
		Version:        o.Version + 1,
	}, nil
}

// IsNeededToUpdateは更新項目がある場合にtrueを返す
func (t *Task) IsNeededToUpdate(p *gr.UpdateTaskRequestParam) bool {
	if p.Name == *new(string) && p.Details == *new(string) && p.ImportanceName == *new(string) && p.Deadline == nil {
		return false
	}

	var deadlineP *time.Time
	if p.Deadline != nil {
		d := p.Deadline.AsTime()
		deadlineP = &d
	} else {
		deadlineP = nil
	}

	return p.Name != t.Name || p.Details != t.Details || p.ImportanceName != t.ImportanceName || *deadlineP != *t.Deadline
}

func (t *Task) String() string {
	return fmt.Sprintf("Id:%v Name:%v Details:%v ImportanceName:%v Deadline:%v RegisteredAt:%v UpdatedAt:%v Version:%v", t.Id, t.Name, t.Details, t.ImportanceName, t.Deadline, t.RegisteredAt, t.UpdatedAt, t.Version)
}
