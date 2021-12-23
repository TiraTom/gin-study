package record

import (
	"fmt"
	"time"
)

type Task struct {
	Id             string
	Name           string
	Details        string
	ImportanceId   int64
	RegisteredTime time.Time
	Deadline       time.Time
	UpdatedTime    time.Time
	Version        uint
}

func (t *Task) String() string {
	return fmt.Sprintf("Id:%v Name:%v Details:%v ImportanceId:%v Deadline:%v RegisteredTime:%v UpdatedTime:%v Version:%v", t.Id, t.Name, t.Details, t.ImportanceId, t.Deadline, t.RegisteredTime, t.UpdatedTime, t.Version)
}
