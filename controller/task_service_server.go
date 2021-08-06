package controller

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpc "github.com/Tiratom/gin-study/grpc"
	"github.com/google/uuid"
)

type TaskServiceServer struct {
}

func (tss *TaskServiceServer) GetAllTasks(ctx context.Context, emp *emptypb.Empty) (*grpc.Tasks, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &grpc.Tasks{
		Tasks: []*grpc.Task{
			{
				Id:           id.String(),
				Name:         "ダミー",
				Details:      "詳細",
				Importance:   grpc.Importance_HIGH,
				RegisteredAt: nowTimestamp,
				Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
				UpdatedAt:    nowTimestamp,
			},
		},
	}, nil
}

func (tss *TaskServiceServer) GetTasks(ctx context.Context, param *grpc.GetTaskByConditionRequestParam) (*grpc.Tasks, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &grpc.Tasks{
		Tasks: []*grpc.Task{
			{
				Id:           id.String(),
				Name:         "ダミー",
				Details:      "詳細",
				Importance:   grpc.Importance_HIGH,
				RegisteredAt: nowTimestamp,
				Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
				UpdatedAt:    nowTimestamp,
			},
		},
	}, nil
}

func (tss *TaskServiceServer) GetTask(ctx context.Context, param *grpc.GetTaskByIdRequestParam) (*grpc.Task, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &grpc.Task{
		Id:           id.String(),
		Name:         "ダミー",
		Details:      "詳細",
		Importance:   grpc.Importance_HIGH,
		RegisteredAt: nowTimestamp,
		Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
		UpdatedAt:    nowTimestamp,
	}, nil
}

func (tss *TaskServiceServer) CreateTask(ctx context.Context, param *grpc.CreateTaskRequestParam) (*grpc.Task, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &grpc.Task{
		Id:           id.String(),
		Name:         "ダミー",
		Details:      "詳細",
		Importance:   grpc.Importance_HIGH,
		RegisteredAt: nowTimestamp,
		Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
		UpdatedAt:    nowTimestamp,
	}, nil
}

func (tss *TaskServiceServer) UpdateTask(ctx context.Context, param *grpc.UpdateTaskRequestParam) (*grpc.Task, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	nowTimestamp := &timestamppb.Timestamp{Seconds: time.Now().Unix()}

	return &grpc.Task{
		Id:           id.String(),
		Name:         "ダミー",
		Details:      "詳細",
		Importance:   grpc.Importance_HIGH,
		RegisteredAt: nowTimestamp,
		Deadline:     &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
		UpdatedAt:    nowTimestamp,
	}, nil
}

func (tss *TaskServiceServer) DeleteTask(ctx context.Context, param *grpc.DeleteTaskRequestParam) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
