package record

import (
	"time"
)

type Task struct {
	Id           string
	Name         string
	Details      string
	ImportanceId int64
	RegisteredAt time.Time
	Deadline     time.Time
	UpdatedAt    time.Time
	Version      uint
}
