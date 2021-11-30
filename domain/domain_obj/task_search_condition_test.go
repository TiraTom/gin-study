package domain_obj

import (
	"reflect"
	"testing"
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTaskSearchCondition_IsDeadlineIncludedInCondition(t *testing.T) {
	type fields struct {
		Name                  string
		Details               string
		ImportanceName        string
		Deadline              *time.Time
		SearchTypeForDeadline *gr.TimestampCompareBy
	}

	dummyTime := time.Now()
	dummySearchType := gr.TimestampCompareBy_AFTER

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "検索条件に期限日時が含まれている場合",
			fields: fields{
				Name:                  "NAME",
				Details:               "DETAILS",
				ImportanceName:        "IMPORTANCE_NAME",
				Deadline:              &dummyTime,
				SearchTypeForDeadline: &dummySearchType,
			},
			want: true,
		},
		{
			name: "検索条件に期限日時が含まれていない場合",
			fields: fields{
				Name:                  "NAME",
				Details:               "DETAILS",
				ImportanceName:        "IMPORTANCE_NAME",
				Deadline:              &dummyTime,
				SearchTypeForDeadline: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TaskSearchCondition{
				Name:                  tt.fields.Name,
				Details:               tt.fields.Details,
				ImportanceName:        tt.fields.ImportanceName,
				Deadline:              tt.fields.Deadline,
				SearchTypeForDeadline: tt.fields.SearchTypeForDeadline,
			}
			if got := tr.IsDeadlineIncludedInCondition(); got != tt.want {
				t.Errorf("TaskSearchCondition.IsDeadlineIncludedInCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskSearchCondition_AsDeadlineConditionSentence(t *testing.T) {
	type fields struct {
		Name                  string
		Details               string
		ImportanceName        string
		Deadline              *time.Time
		SearchTypeForDeadline *gr.TimestampCompareBy
	}

	dummyTime := time.Now()
	searchTypeAfter := gr.TimestampCompareBy_AFTER
	searchTypeBefore := gr.TimestampCompareBy_BEFORE
	searchTypeSame := gr.TimestampCompareBy_SAME

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "指定した日時より前の期限日時のタスクを探す場合",
			fields: fields{
				Name:                  "DUMMY_NAME",
				Details:               "DUMMY_DETAILS",
				ImportanceName:        "DUMMY_IMPORTANCE_NAME",
				Deadline:              &dummyTime,
				SearchTypeForDeadline: &searchTypeBefore,
			},
			want:    "tasks.deadline < ?",
			wantErr: false,
		},
		{
			name: "指定した日時より後の期限日時のタスクを探す場合",
			fields: fields{
				Name:                  "DUMMY_NAME",
				Details:               "DUMMY_DETAILS",
				ImportanceName:        "DUMMY_IMPORTANCE_NAME",
				Deadline:              &dummyTime,
				SearchTypeForDeadline: &searchTypeAfter,
			},
			want:    "tasks.deadline > ?",
			wantErr: false,
		},
		{
			name: "指定した日時と同じ期限日時のタスクを探す場合",
			fields: fields{
				Name:                  "DUMMY_NAME",
				Details:               "DUMMY_DETAILS",
				ImportanceName:        "DUMMY_IMPORTANCE_NAME",
				Deadline:              &dummyTime,
				SearchTypeForDeadline: &searchTypeSame,
			},
			want:    "tasks.deadline = ?",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TaskSearchCondition{
				Name:                  tt.fields.Name,
				Details:               tt.fields.Details,
				ImportanceName:        tt.fields.ImportanceName,
				Deadline:              tt.fields.Deadline,
				SearchTypeForDeadline: tt.fields.SearchTypeForDeadline,
			}
			got, err := tr.AsDeadlineConditionSentence()
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskSearchCondition.AsDeadlineConditionSentence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TaskSearchCondition.AsDeadlineConditionSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskSearchCondition_AsSelectConditionMap(t *testing.T) {
	type fields struct {
		Name                  string
		Details               string
		ImportanceName        string
		Deadline              *time.Time
		SearchTypeForDeadline *gr.TimestampCompareBy
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "name,details,importanceNameの全て値が存在する場合",
			fields: fields{
				Name:                  "DUMMY_NAME",
				Details:               "DUMMY_DETAILS",
				ImportanceName:        "DUMMY_IMPORTANCE_NAME",
				Deadline:              nil,
				SearchTypeForDeadline: nil,
			},
			want: map[string]interface{}{
				"tasks.name":       "DUMMY_NAME",
				"tasks.details":    "DUMMY_DETAILS",
				"importances.name": "DUMMY_IMPORTANCE_NAME",
			},
		},
		{
			name: "name,details,importanceNameの全て値が存在しない場合",
			fields: fields{
				Name:                  "",
				Details:               "",
				ImportanceName:        "",
				Deadline:              nil,
				SearchTypeForDeadline: nil,
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TaskSearchCondition{
				Name:                  tt.fields.Name,
				Details:               tt.fields.Details,
				ImportanceName:        tt.fields.ImportanceName,
				Deadline:              tt.fields.Deadline,
				SearchTypeForDeadline: tt.fields.SearchTypeForDeadline,
			}
			if got := tr.AsSelectConditionMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskSearchCondition.AsSelectConditionMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTaskSearchCondition(t *testing.T) {
	type args struct {
		p *gr.GetTaskByConditionRequestParam
	}

	// 2021/11/29 00:11:12(UTC) のタイムスタンプ
	dummyTimeStampSec := int64(1638144672)
	dummyTime, err := time.Parse("2006/01/02 15:04:05", "2021/11/29 00:11:12")
	if err != nil {
		t.Errorf("時刻の返還処理でエラー発生 %w", err)
	}

	dummySearchTypeForDeadline := gr.TimestampCompareBy_AFTER

	tests := []struct {
		name string
		args args
		want *TaskSearchCondition
	}{
		{
			name: "すべてに値が設定されている場合",
			args: args{
				p: &gr.GetTaskByConditionRequestParam{
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "DUMMY_IMPORTANCE_NAME",
					Deadline: &timestamppb.Timestamp{
						Seconds: dummyTimeStampSec,
					},
					SearchTypeForDeadline: dummySearchTypeForDeadline,
				},
			},
			want: &TaskSearchCondition{
				Name:                  "DUMMY_NAME",
				Details:               "DUMMY_DETAILS",
				ImportanceName:        "DUMMY_IMPORTANCE_NAME",
				Deadline:              &dummyTime,
				SearchTypeForDeadline: &dummySearchTypeForDeadline,
			},
		},
		{
			name: "値が全く設定されていない場合",
			args: args{
				p: &gr.GetTaskByConditionRequestParam{},
			},
			want: &TaskSearchCondition{
				Name:                  "",
				Details:               "",
				ImportanceName:        "",
				Deadline:              nil,
				SearchTypeForDeadline: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskSearchCondition(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskSearchCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}
