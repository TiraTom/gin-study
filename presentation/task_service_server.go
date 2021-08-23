package presentation

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/usecase"
	"github.com/google/uuid"
)

type TaskServiceServer struct {
	log *middleware.ZapLogger
	gtu *usecase.GetTask
}

func (tss *TaskServiceServer) GetAllTasks(ctx context.Context, emp *emptypb.Empty) (*gr.Tasks, error) {

	tss.log.Info(ctx, "HOGEHOGE")

	allTasks, err := tss.gtu.GetAllTasks()
	if err != nil {
		return nil, err
	}

	return allTasks, nil
}

func (tss *TaskServiceServer) GetTasks(ctx context.Context, param *gr.GetTaskByConditionRequestParam) (*gr.Tasks, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &gr.Tasks{
		Tasks: []*gr.Task{
			{
				Id:           id.String(),
				Name:         "ダミー",
				Details:      "詳細",
				Importance:   gr.Importance_HIGH,
				RegisteredAt: nowTimestamp,
				Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
				UpdatedAt:    nowTimestamp,
			},
		},
	}, nil
}

func (tss *TaskServiceServer) GetTask(ctx context.Context, param *gr.GetTaskByIdRequestParam) (*gr.Task, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &gr.Task{
		Id:           id.String(),
		Name:         "ダミー",
		Details:      "詳細",
		Importance:   gr.Importance_HIGH,
		RegisteredAt: nowTimestamp,
		Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
		UpdatedAt:    nowTimestamp,
	}, nil
}

func (tss *TaskServiceServer) CreateTask(ctx context.Context, param *gr.CreateTaskRequestParam) (*gr.Task, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &gr.Task{
		Id:           id.String(),
		Name:         "ダミー",
		Details:      "詳細",
		Importance:   gr.Importance_HIGH,
		RegisteredAt: nowTimestamp,
		Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
		UpdatedAt:    nowTimestamp,
	}, nil
}

func (tss *TaskServiceServer) UpdateTask(ctx context.Context, param *gr.UpdateTaskRequestParam) (*gr.Task, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &gr.Task{
		Id:           id.String(),
		Name:         "ダミー",
		Details:      "詳細",
		Importance:   gr.Importance_HIGH,
		RegisteredAt: nowTimestamp,
		Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
		UpdatedAt:    nowTimestamp,
	}, nil
}

func (tss *TaskServiceServer) DeleteTask(ctx context.Context, param *gr.DeleteTaskRequestParam) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func NewTaskServiceServer(log *middleware.ZapLogger, gtu *usecase.GetTask) *TaskServiceServer {
	return &TaskServiceServer{log, gtu}
}
