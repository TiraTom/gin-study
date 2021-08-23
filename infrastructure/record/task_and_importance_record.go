package infrastructure

import (
	"time"
)

type TaskAndImportanceRecord struct {
	Id              string
	Name            string
	Details         string
	ImportanceId    int
	ImportanceName  string
	ImportanceLevel int
	RegisteredAt    time.Time
	Deadline        time.Time
	UpdatedAt       time.Time
}
