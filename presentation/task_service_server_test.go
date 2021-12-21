package presentation_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/presentation"
	"github.com/Tiratom/gin-study/usecase/usecase_interface"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		{
			name: "タスクが存在する場合",
			mock: &GetTaskUsecaseMock{
				FakeDoAll: func() (*domain_obj.Tasks, error) {
					return &domain_obj.Tasks{
						Value: []*domain_obj.Task{
							{
								Id:             "DUMMY_ID",
								Name:           "DUMMY_NAME",
								Details:        "DUMMY_DETAILS",
								ImportanceName: "DUMMY_IMPORTANCE_NAME",
								Deadline:       &time20210822150001,
								RegisteredAt:   &time20210822150002,
								UpdatedAt:      &time20210822150003,
								Version:        2,
							},
						},
					}, nil
				},
			},
			want: &gr.Tasks{
				Tasks: []*gr.Task{
					{
						Id:             "DUMMY_ID",
						Name:           "DUMMY_NAME",
						Details:        "DUMMY_DETAILS",
						ImportanceName: "DUMMY_IMPORTANCE_NAME",
						Deadline:       &timestamppb.Timestamp{Seconds: timestamp20210822150001},
						RegisteredAt:   &timestamppb.Timestamp{Seconds: timestamp20210822150002},
						UpdatedAt:      &timestamppb.Timestamp{Seconds: timestamp20210822150003},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "usecase層の処理でエラーが発生した場合panicが起きずにエラーを返却する",
			mock: &GetTaskUsecaseMock{
				FakeDoAll: func() (*domain_obj.Tasks, error) {
					return nil, fmt.Errorf("usecase層でのエラー")
				},
			},
			want:    nil,
			wantErr: true,
		},
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
