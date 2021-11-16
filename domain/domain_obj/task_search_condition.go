package domain_obj

import (
	"fmt"
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

func (t *TaskSearchCondition) IsDeadlineIncludedInCondition() bool {
	return t.SearchTypeForDeadline != nil
}

func (t *TaskSearchCondition) AsDeadlineConditionSentence() (string, error) {
	switch t.SearchTypeForDeadline.Number() {
	case gr.TimestampCompareBy_BEFORE.Number():
		return "tasks.deadline < ?", nil
	case gr.TimestampCompareBy_SAME.Number():
		return "tasks.deadline = ?", nil
	case gr.TimestampCompareBy_AFTER.Number():
		return "tasks.deadline > ?", nil
	default:
		return "", fmt.Errorf("期限日時の比較条件設定として使えない値です %s", t.SearchTypeForDeadline)
	}
}

func (t *TaskSearchCondition) AsSelectConditionMap() map[string]interface{} {
	c := make(map[string]interface{}, 3)

	if t.Name != "" {
		c["tasks.name"] = t.Name
	}
	if t.Details != "" {
		c["tasks.details"] = t.Details
	}
	if t.ImportanceName != "" {
		c["importances.name"] = t.ImportanceName
	}

	return c
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
