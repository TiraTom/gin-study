package presentation

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/middleware"
	"github.com/Tiratom/gin-study/usecase/usecase_interface"
)

type TaskServiceServer struct {
	Log               *middleware.ZapLogger
	GetTaskUsecase    usecase_interface.GetTask
	CreateTaskUsecase usecase_interface.CreateTask
	UpdateTaskUsecase usecase_interface.UpdateTask
	DeleteTaskUsecase usecase_interface.DeleteTask
	SearchTaskUsecase usecase_interface.SearchTask
}

func (tss *TaskServiceServer) GetAllTasks(ctx context.Context, emp *emptypb.Empty) (*gr.Tasks, error) {
	tasks, err := tss.GetTaskUsecase.DoAll()
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました; %w", err)
	}

	ts, err := tasks.ToDto()
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("タスク全件取得後の変換処理においてエラーが発生しました; %w", err)
	}

	return ts, err
}

func (tss *TaskServiceServer) GetTasks(ctx context.Context, param *gr.GetTaskByConditionRequestParam) (*gr.Tasks, error) {
	// TODO ageとかゼロ値がありうる項目も検索する場合にも問題ないようにポインタでparamも設定すべきか？
	tasks, err := tss.SearchTaskUsecase.Do(param)
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
	}

	t, err := tasks.ToDto()
	if err != nil {
		return nil, fmt.Errorf("タスク検索結果の変換処理においてエラーが発生しました(検索結果=%v); %v", tasks, err)
	}

	return t, err
}

func (tss *TaskServiceServer) GetTask(ctx context.Context, param *gr.GetTaskByIdRequestParam) (*gr.Task, error) {
	task, err := tss.GetTaskUsecase.DoById(param.Id)
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました; %w", err)
	}

	t, err := task.ToDto()
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("タスク取得後の変換処理においてエラーが発生しました; %w", err)
	}

	return t, err
}

func (tss *TaskServiceServer) CreateTask(ctx context.Context, param *gr.CreateTaskRequestParam) (*gr.Task, error) {
	task, err := tss.CreateTaskUsecase.Do(param)
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました; %w", err)
	}

	t, err := task.ToDto()
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("タスク作成成功後戻り値生成においてエラーが発生しました(p={Name=%v Details=%v ImportanceName=%v Deadline=%v}); %w", t.Name, t.Details, t.ImportanceName, t.Deadline, err)
	}

	return t, err
}

func (tss *TaskServiceServer) UpdateTask(ctx context.Context, param *gr.UpdateTaskRequestParam) (*gr.Task, error) {
	// ※idと更新内容のパラメーターを別の変数として渡してもらおうと思ったが、protoファイルの定義上引数は１つにするのが定石っぽく（変更に強くするための模様）paramの中にidも変更内容も持たせてある

	updateResult, err := tss.UpdateTaskUsecase.Do(param)
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました; %w", err)
	}

	t, err := updateResult.ToDto()
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("タスク更新結果の変換処理においてエラーが発生しました(%v)); %w", updateResult, err)
	}

	return t, nil
}

func (tss *TaskServiceServer) DeleteTask(ctx context.Context, param *gr.DeleteTaskRequestParam) (*emptypb.Empty, error) {
	err := tss.DeleteTaskUsecase.Do(param)
	if err != nil {
		tss.Log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました; ", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました); %w", err)
	}

	return &emptypb.Empty{}, err
}

func NewTaskServiceServer(log *middleware.ZapLogger, gtu usecase_interface.GetTask, ctu usecase_interface.CreateTask, utu usecase_interface.UpdateTask, dtu usecase_interface.DeleteTask, stu usecase_interface.SearchTask) *TaskServiceServer {
	return &TaskServiceServer{log, gtu, ctu, utu, dtu, stu}
}
