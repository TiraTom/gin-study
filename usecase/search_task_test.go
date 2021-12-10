package usecase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/mock/mock_repository_interface"
	"github.com/golang/mock/gomock"
)

func TestSearchTask_Do(t *testing.T) {
	type args struct {
		p *gr.GetTaskByConditionRequestParam
	}
	tests := []struct {
		name            string
		args            args
		want            *domain_obj.Tasks
		wantErr         bool
		prepareMockFunc func() // mock準備用関数
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mtaskR := mock_repository_interface.NewMockTask(ctrl)
			mtaskR.EXPECT().Search(tt.args).DoAndReturn(func() (*domain_obj.Tasks, error) {
				return nil, fmt.Errorf("not yet implemented")
			})

			s := &SearchTask{
				tr: mtaskR,
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
