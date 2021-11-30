package domain_obj

import (
	"reflect"
	"testing"
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTasks_ToDto(t *testing.T) {
	type fields struct {
		Value []*Task
	}

	dummyTime1, err := time.Parse("2006/01/02 15:04:05", "2021/11/29 00:11:12")
	if err != nil {
		t.Errorf("時刻の変換処理でエラー発生 %w", err)
	}
	// 2021/11/29 00:11:12(UTC)
	dummyTimeStampSec1 := int64(1638144672)

	// 2021/11/30 00:11:12(UTC)
	dummyTime2 := dummyTime1.AddDate(0, 0, 1)
	dummyTimeStampSec2 := int64(1638231072)

	// 2021/11/30 00:11:12(UTC)
	dummyTime3 := dummyTime2.AddDate(0, 0, 1)
	dummyTimeStampSec3 := int64(1638317472)

	tests := []struct {
		name    string
		fields  fields
		want    *gr.Tasks
		wantErr bool
	}{
		{
			name: "通常パターン",
			fields: fields{
				Value: []*Task{
					{
						Id:             "DUMMY_ID1",
						Name:           "DUMMY_NAME1",
						Details:        "DUMMY_DETAILS1",
						ImportanceName: "DUMMY_IMPORTANCE_NAME1",
						RegisteredAt:   &dummyTime1,
						Deadline:       &dummyTime2,
						UpdatedAt:      &dummyTime3,
						Version:        2,
					},
					{
						Id:             "DUMMY_ID2",
						Name:           "DUMMY_NAME2",
						Details:        "DUMMY_DETAILS2",
						ImportanceName: "DUMMY_IMPORTANCE_NAME2",
						RegisteredAt:   &dummyTime1,
						Deadline:       &dummyTime2,
						UpdatedAt:      &dummyTime3,
						Version:        3,
					},
				},
			},
			want: &gr.Tasks{
				Tasks: []*gr.Task{
					{
						Id:             "DUMMY_ID1",
						Name:           "DUMMY_NAME1",
						Details:        "DUMMY_DETAILS1",
						ImportanceName: "DUMMY_IMPORTANCE_NAME1",
						RegisteredAt: &timestamppb.Timestamp{
							Seconds: dummyTimeStampSec1,
						},
						Deadline: &timestamppb.Timestamp{
							Seconds: dummyTimeStampSec2,
						},
						UpdatedAt: &timestamppb.Timestamp{
							Seconds: dummyTimeStampSec3,
						},
					},
					{
						Id:             "DUMMY_ID2",
						Name:           "DUMMY_NAME2",
						Details:        "DUMMY_DETAILS2",
						ImportanceName: "DUMMY_IMPORTANCE_NAME2",
						RegisteredAt: &timestamppb.Timestamp{
							Seconds: dummyTimeStampSec1,
						},
						Deadline: &timestamppb.Timestamp{
							Seconds: dummyTimeStampSec2,
						},
						UpdatedAt: &timestamppb.Timestamp{
							Seconds: dummyTimeStampSec3,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "空の値があるパターン",
			fields: fields{
				Value: []*Task{
					{
						Id:             "DUMMY_ID1",
						Name:           "DUMMY_NAME1",
						Details:        "DUMMY_DETAILS1",
						ImportanceName: "DUMMY_IMPORTANCE_NAME1",
						RegisteredAt:   nil,
						Deadline:       nil,
						UpdatedAt:      nil,
						Version:        2,
					},
				},
			},
			want: &gr.Tasks{
				Tasks: []*gr.Task{
					{
						Id:             "DUMMY_ID1",
						Name:           "DUMMY_NAME1",
						Details:        "DUMMY_DETAILS1",
						ImportanceName: "DUMMY_IMPORTANCE_NAME1",
						RegisteredAt:   nil,
						Deadline:       nil,
						UpdatedAt:      nil,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tasks{
				Value: tt.fields.Value,
			}
			got, err := tr.ToDto()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tasks.ToDto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tasks.ToDto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTasks(t *testing.T) {
	type args struct {
		records []*infrastructure.TaskAndImportance
	}

	// 2021/11/29 00:11:12(UTC)
	dummyTime1, err := time.Parse("2006/01/02 15:04:05", "2021/11/29 00:11:12")
	if err != nil {
		t.Errorf("時刻の変換処理でエラー発生 %w", err)
	}
	// 2021/11/30 00:11:12(UTC)
	dummyTime2 := dummyTime1.AddDate(0, 0, 1)
	// 2021/11/30 00:11:12(UTC)
	dummyTime3 := dummyTime2.AddDate(0, 0, 1)

	tests := []struct {
		name string
		args args
		want *Tasks
	}{
		{
			name: "通常パターン",
			args: args{
				records: []*infrastructure.TaskAndImportance{
					{
						Id:              "DUMMY_ID1",
						Name:            "DUMMY_NAME1",
						Details:         "DUMMY_DETAILS1",
						ImportanceId:    2,
						ImportanceName:  "DUMMY_IMPORTANCE_NAME1",
						ImportanceLevel: 3,
						RegisteredAt:    dummyTime1,
						Deadline:        dummyTime2,
						UpdatedAt:       dummyTime3,
						Version:         4,
					},
					{
						Id:              "DUMMY_ID2",
						Name:            "DUMMY_NAME2",
						Details:         "DUMMY_DETAILS2",
						ImportanceId:    12,
						ImportanceName:  "DUMMY_IMPORTANCE_NAME2",
						ImportanceLevel: 13,
						RegisteredAt:    dummyTime1,
						Deadline:        dummyTime2,
						UpdatedAt:       dummyTime3,
						Version:         14,
					},
				},
			},
			want: &Tasks{
				Value: []*Task{
					{
						Id:             "DUMMY_ID1",
						Name:           "DUMMY_NAME1",
						Details:        "DUMMY_DETAILS1",
						ImportanceName: "DUMMY_IMPORTANCE_NAME1",
						RegisteredAt:   &dummyTime1,
						Deadline:       &dummyTime2,
						UpdatedAt:      &dummyTime3,
						Version:        4,
					},
					{
						Id:             "DUMMY_ID2",
						Name:           "DUMMY_NAME2",
						Details:        "DUMMY_DETAILS2",
						ImportanceName: "DUMMY_IMPORTANCE_NAME2",
						RegisteredAt:   &dummyTime1,
						Deadline:       &dummyTime2,
						UpdatedAt:      &dummyTime3,
						Version:        14,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTasks(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
