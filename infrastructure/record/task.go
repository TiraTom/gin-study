package record

import (
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
