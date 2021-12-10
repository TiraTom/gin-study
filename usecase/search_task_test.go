package usecase

import (
	"fmt"
	"reflect"
	"testing"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/mock/mock_repository_interface"
	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSearchTask_Do(t *testing.T) {
	type args struct {
		p *gr.GetTaskByConditionRequestParam
	}
	tests := []struct {
		name            string
		args            args
		want            *gr.Tasks
		wantErr         bool
		prepareMockFunc func() // mock準備用関数
	}{
		{
			name: "",
			args: args{
				p: &gr.GetTaskByConditionRequestParam{
					Name:                  "",
					Details:               "",
					ImportanceName:        "",
					Deadline:              &timestamppb.Timestamp{},
					SearchTypeForDeadline: 0,
				},
			},
			want: &gr.Tasks{
				Tasks: []*gr.Task{},
			},
			wantErr: false,
			prepareMockFunc: func() {
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mTask := mock_repository_interface.NewMockTask(ctrl)
			mTask.EXPECT().Search(tt.args).Return(func() (*gr.Tasks, error) {
				return nil, fmt.Errorf("not yet implemented")
			})

			s := &SearchTask{
				tr: mTask,
			}
			got, err := s.Do(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchTask.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchTask.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
