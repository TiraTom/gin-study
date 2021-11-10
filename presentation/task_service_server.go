package presentation

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/usecase"
	"github.com/google/uuid"
)

type TaskServiceServer struct {
	log        *middleware.ZapLogger
	getTask    *usecase.GetTask
	createTask *usecase.CreateTask
	updateTask *usecase.UpdateTask
	deleteTask *usecase.DeleteTask
}

func (tss *TaskServiceServer) GetAllTasks(ctx context.Context, emp *emptypb.Empty) (*gr.Tasks, error) {

	tss.log.Info(ctx, "HOGEHOGE")

	allTasks, err := tss.getTask.GetAllTasks()
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
				Id:             id.String(),
				Name:           "ダミー",
				Details:        "詳細",
				ImportanceName: "HIGH",
				RegisteredAt:   nowTimestamp,
				Deadline:       &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
				UpdatedAt:      nowTimestamp,
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
		Id:             id.String(),
		Name:           "ダミー",
		Details:        "詳細",
		ImportanceName: "HIGH",
		RegisteredAt:   nowTimestamp,
		Deadline:       &timestamppb.Timestamp{Seconds: time.Date(2021, 8, 6, 12, 0, 0, 0, tokyo).Unix()},
		UpdatedAt:      nowTimestamp,
	}, nil
}

func (tss *TaskServiceServer) CreateTask(ctx context.Context, param *gr.CreateTaskRequestParam) (*gr.Task, error) {
	return tss.createTask.Do(param)
}

func (tss *TaskServiceServer) UpdateTask(ctx context.Context, param *gr.UpdateTaskRequestParam) (*gr.Task, error) {
	// ※idと更新内容のパラメーターを別の変数として渡してもらおうと思ったが、protoファイルの定義上引数は１つにするのが定石っぽく（変更に強くするための模様）paramの中にidも変更内容も持たせてある

	t, err := tss.updateTask.Do(param)
	if err != nil {
		return nil, err
	}

	// memo: 返却用データに詰め替えるのをどこでやるべきかは悩み中、、
	updatedTask, err := t.ToDto()
	if err != nil {
		return nil, fmt.Errorf("データ更新成功後、内部エラーが発生しました %w", err)
	}

	return updatedTask, nil
}

func (tss *TaskServiceServer) DeleteTask(ctx context.Context, param *gr.DeleteTaskRequestParam) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, tss.deleteTask.Do(param)
}

func NewTaskServiceServer(log *middleware.ZapLogger, gtu *usecase.GetTask, ctu *usecase.CreateTask, utu *usecase.UpdateTask, dtu *usecase.DeleteTask) *TaskServiceServer {
	return &TaskServiceServer{log, gtu, ctu, utu, dtu}
}
