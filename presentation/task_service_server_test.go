package presentation_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	"github.com/Tiratom/gin-study/usecase/usecase_interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

// 勉強がてらなのでpresentation層のテストはGetAllTasksの分だけ書いて終わりにする

func TestTaskServiceServer_GetAllTasks(t *testing.T) {
	tests := []struct {
		name    string
		mock    *GetTaskUsecaseMock
		want    *gr.Tasks
		wantErr bool
	}{
		{
			name: "タスクが空の場合",
			mock: &GetTaskUsecaseMock{
				FakeDoAll: func() (*domain_obj.Tasks, error) {
					return &domain_obj.Tasks{}, nil
				},
			},
			want: &gr.Tasks{
				Tasks: []*gr.Task{},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tss := &presentation.TaskServiceServer{
				Log:               &middleware.ZapLogger{},
				GetTaskUsecase:    tt.mock,
				CreateTaskUsecase: &CreateTaskUsecaseEmptyMock{},
				UpdateTaskUsecase: &UpdateTaskUsecaseEmptyMock{},
				DeleteTaskUsecase: &DeleteTaskUsecaseEmptyMock{},
				SearchTaskUsecase: &SearchTaskUsecaseEmptyMock{},
			}
			got, err := tss.GetAllTasks(context.Background(), &emptypb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskServiceServer.GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskServiceServer.GetAllTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

type GetTaskUsecaseMock struct {
	FakeDoAll  func() (*domain_obj.Tasks, error)
	FakeDoById func(id string) (*domain_obj.Task, error)
}

func (m *GetTaskUsecaseMock) DoAll() (*domain_obj.Tasks, error) {
	return m.FakeDoAll()
}

func (m *GetTaskUsecaseMock) DoById(id string) (*domain_obj.Task, error) {
	return m.FakeDoById(id)
}

type CreateTaskUsecaseEmptyMock struct {
	usecase_interface.CreateTask
}
type UpdateTaskUsecaseEmptyMock struct {
	usecase_interface.UpdateTask
}
type DeleteTaskUsecaseEmptyMock struct {
	usecase_interface.DeleteTask
}
type SearchTaskUsecaseEmptyMock struct {
	usecase_interface.SearchTask
}
