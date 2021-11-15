package domain_obj

import (
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
)

type TaskSearchCondition struct {
	Name                  string
	Details               string
	ImportanceName        string
	Deadline              *time.Time
	SearchTypeForDeadline *gr.TimestampCompareBy
}

func (t *TaskSearchCondition) IsImportanceIncludedInCondition() bool {
	return t.ImportanceName != ""
}

func (t *TaskSearchCondition) IsDeadlineIncludedInCondition() bool {
	return t.SearchTypeForDeadline != nil
}

func NewTaskSearchCondition(p *gr.GetTaskByConditionRequestParam) *TaskSearchCondition {
	dl := p.Deadline.AsTime()

	return &TaskSearchCondition{
		Name:                  p.Name,
		Details:               p.Details,
		ImportanceName:        p.ImportanceName,
		Deadline:              &dl,
		SearchTypeForDeadline: &p.SearchTypeForDeadline,
	}
}
