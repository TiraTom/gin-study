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
	tasks, err := tss.getTask.DoAll()
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err.Error()))
		return nil, fmt.Errorf("処理中にエラーが発生しました: %w", err)
	}

	ts, err := tasks.ToDto()
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err.Error()))
		return nil, fmt.Errorf("タスク全件取得後の変換処理においてエラーが発生しました: %w", err)
	}

	return ts, err
}

func (tss *TaskServiceServer) GetTasks(ctx context.Context, param *gr.GetTaskByConditionRequestParam) (*gr.Tasks, error) {
	// TODO ageとかゼロ値がありうる項目も検索する場合にも問題ないようにポインタでparamも設定すべきか？
	tasks, err := tss.searchTask.Do(param)
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err))
	}

	t, err := tasks.ToDto()
	if err != nil {
		return nil, fmt.Errorf("タスク検索結果の変換処理においてエラーが発生しました(検索結果=%v):%w", tasks, err)
	}

	return t, err
}

func (tss *TaskServiceServer) GetTask(ctx context.Context, param *gr.GetTaskByIdRequestParam) (*gr.Task, error) {
	task, err := tss.getTask.DoById(param.Id)
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました: %w", err)
	}

	t, err := task.ToDto()
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err.Error()))
		return nil, fmt.Errorf("タスク取得後の変換処理においてエラーが発生しました: %w", err)
	}

	return t, err
}

func (tss *TaskServiceServer) CreateTask(ctx context.Context, param *gr.CreateTaskRequestParam) (*gr.Task, error) {
	task, err := tss.createTask.Do(param)
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました: %w", err)
	}

	t, err := task.ToDto()
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err.Error()))
		return nil, fmt.Errorf("タスク作成成功後戻り値生成においてエラーが発生しました(p={Name=%v Details=%v ImportanceName=%v Deadline=%v}): %w", t.Name, t.Details, t.ImportanceName, t.Deadline, err)
	}

	return t, err
}

func (tss *TaskServiceServer) UpdateTask(ctx context.Context, param *gr.UpdateTaskRequestParam) (*gr.Task, error) {
	// ※idと更新内容のパラメーターを別の変数として渡してもらおうと思ったが、protoファイルの定義上引数は１つにするのが定石っぽく（変更に強くするための模様）paramの中にidも変更内容も持たせてある

	updateResult, err := tss.updateTask.Do(param)
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err))
		return nil, fmt.Errorf("処理中にエラーが発生しました: %w", err)
	}

	t, err := updateResult.ToDto()
	if err != nil {
		tss.log.Warn(ctx, fmt.Sprint("処理中にエラーが発生しました: %w", err))
		return nil, fmt.Errorf("タスク更新結果の変換処理においてエラーが発生しました(%v)):%w", updateResult, err)
	}

	return t, nil
}

func (tss *TaskServiceServer) DeleteTask(ctx context.Context, param *gr.DeleteTaskRequestParam) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, tss.deleteTask.Do(param)
}

func NewTaskServiceServer(log *middleware.ZapLogger, gtu *usecase.GetTask, ctu *usecase.CreateTask, utu *usecase.UpdateTask, dtu *usecase.DeleteTask, stu *usecase.SearchTask) *TaskServiceServer {
	return &TaskServiceServer{log, gtu, ctu, utu, dtu, stu}
}
