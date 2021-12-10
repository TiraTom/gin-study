package usecase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/domain/repository_interface"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/mock/mock_repository_interface"
	"github.com/golang/mock/gomock"
)

func TestGetTask_DoAll(t *testing.T) {
	tests := []struct {
		name            string
		want            *gr.Tasks
		wantErr         bool
		prepareMockFunc func() // mockテスト用関数
	}{
		{
			name:    "",
			want:    &gr.Tasks{},
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
			mTask.EXPECT().GetAll().DoAndReturn(func() (*domain_obj.Tasks, error) {
				return nil, fmt.Errorf("not yet implemented")
			})

			gt := &GetTask{
				tr: mTask,
			}
			got, err := gt.DoAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask.DoAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask.DoAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTask_DoById(t *testing.T) {
	type fields struct {
		tr repository_interface.Task
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gr.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gt := &GetTask{
				tr: tt.fields.tr,
			}
			got, err := gt.DoById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask.DoById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask.DoById() = %v, want %v", got, tt.want)
			}
		})
	}
}
