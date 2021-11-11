package presentation

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/usecase"
)

type TaskServiceServer struct {
	log        *middleware.ZapLogger
	getTask    *usecase.GetTask
	createTask *usecase.CreateTask
	updateTask *usecase.UpdateTask
	deleteTask *usecase.DeleteTask
	searchTask *usecase.SearchTask
}

func (tss *TaskServiceServer) GetAllTasks(ctx context.Context, emp *emptypb.Empty) (*gr.Tasks, error) {
	return tss.getTask.DoAll()
}

func (tss *TaskServiceServer) GetTasks(ctx context.Context, param *gr.GetTaskByConditionRequestParam) (*gr.Tasks, error) {
	return tss.searchTask.Do(param)
}

func (tss *TaskServiceServer) GetTask(ctx context.Context, param *gr.GetTaskByIdRequestParam) (*gr.Task, error) {
	return tss.getTask.DoById(param.Id)
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

func NewTaskServiceServer(log *middleware.ZapLogger, gtu *usecase.GetTask, ctu *usecase.CreateTask, utu *usecase.UpdateTask, dtu *usecase.DeleteTask, stu *usecase.SearchTask) *TaskServiceServer {
	return &TaskServiceServer{log, gtu, ctu, utu, dtu, stu}
}
