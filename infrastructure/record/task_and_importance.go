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
	RegisteredAt    time.Time
	Deadline        time.Time
	UpdatedAt       time.Time
	Version         uint
}
