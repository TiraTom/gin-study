package record

import (
	"time"
)

type TaskAndImportance struct {
	Id              string
	Name            string
	Details         string
	ImportanceId    int
	ImportanceName  string
	ImportanceLevel int
	RegisteredTime  time.Time
	Deadline        time.Time
	UpdatedTime     time.Time
	Version         uint
}
