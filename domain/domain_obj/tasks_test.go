package domain_obj

import (
	"reflect"
	"testing"

	gr "github.com/Tiratom/gin-study/grpc"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTasks_ToDto(t *testing.T) {
	type fields struct {
		Value []*Task
	}

	testVs := getDummyValues(t)

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
						RegisteredAt:   &testVs.dummyTime1,
						Deadline:       &testVs.dummyTime2,
						UpdatedAt:      &testVs.dummyTime3,
						Version:        2,
					},
					{
						Id:             "DUMMY_ID2",
						Name:           "DUMMY_NAME2",
						Details:        "DUMMY_DETAILS2",
						ImportanceName: "DUMMY_IMPORTANCE_NAME2",
						RegisteredAt:   &testVs.dummyTime1,
						Deadline:       &testVs.dummyTime2,
						UpdatedAt:      &testVs.dummyTime3,
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
							Seconds: testVs.dummyTimestampSec1,
						},
						Deadline: &timestamppb.Timestamp{
							Seconds: testVs.dummyTimestampSec2,
						},
						UpdatedAt: &timestamppb.Timestamp{
							Seconds: testVs.dummyTimestampSec3,
						},
					},
					{
						Id:             "DUMMY_ID2",
						Name:           "DUMMY_NAME2",
						Details:        "DUMMY_DETAILS2",
						ImportanceName: "DUMMY_IMPORTANCE_NAME2",
						RegisteredAt: &timestamppb.Timestamp{
							Seconds: testVs.dummyTimestampSec1,
						},
						Deadline: &timestamppb.Timestamp{
							Seconds: testVs.dummyTimestampSec2,
						},
						UpdatedAt: &timestamppb.Timestamp{
							Seconds: testVs.dummyTimestampSec3,
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

	testVs := getDummyValues(t)

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
						RegisteredTime:  testVs.dummyTime1,
						Deadline:        testVs.dummyTime2,
						UpdatedTime:     testVs.dummyTime3,
						Version:         4,
					},
					{
						Id:              "DUMMY_ID2",
						Name:            "DUMMY_NAME2",
						Details:         "DUMMY_DETAILS2",
						ImportanceId:    12,
						ImportanceName:  "DUMMY_IMPORTANCE_NAME2",
						ImportanceLevel: 13,
						RegisteredTime:  testVs.dummyTime1,
						Deadline:        testVs.dummyTime2,
						UpdatedTime:     testVs.dummyTime3,
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
						RegisteredAt:   &testVs.dummyTime1,
						Deadline:       &testVs.dummyTime2,
						UpdatedAt:      &testVs.dummyTime3,
						Version:        4,
					},
					{
						Id:             "DUMMY_ID2",
						Name:           "DUMMY_NAME2",
						Details:        "DUMMY_DETAILS2",
						ImportanceName: "DUMMY_IMPORTANCE_NAME2",
						RegisteredAt:   &testVs.dummyTime1,
						Deadline:       &testVs.dummyTime2,
						UpdatedAt:      &testVs.dummyTime3,
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
