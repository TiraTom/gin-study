package repository_impl

import (
	"fmt"

	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/infrastructure/record"
)

type Task struct {
	db *config.DB
}

func (t *Task) GetAll() []*domain_obj.Task {
	var records []*record.TaskAndImportance
	// TODO この書き方だとimportance系の値を入れてくれない理由がわからない（selectの並び順に決まりがある？）
	// t.db.Gdb.Table("tasks").Select("tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_at as registered_at", "tasks.deadline as deadline", "tasks.updated_at as updated_at", "importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Scan(&records)
	t.db.Gdb.Table("tasks").Select("importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level", "tasks.version as version", "tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_at as registered_at", "tasks.deadline as deadline", "tasks.updated_at as updated_at").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Scan(&records)

	result := make([]*domain_obj.Task, len(records))
	for i, v := range records {
		result[i] = domain_obj.NewTask(v)
	}

	return result
}

func (t *Task) GetById(id string) (*domain_obj.Task, error) {
	var foundTask *record.TaskAndImportance
	result := t.db.Gdb.Table("tasks").Select("importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level", "tasks.version as version", "tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_at as registered_at", "tasks.deadline as deadline", "tasks.updated_at as updated_at").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Where("tasks.id = ?", id).First(&foundTask)

	if result.Error != nil {
		return nil, fmt.Errorf("該当のタスクは存在しません")
	}

	return domain_obj.NewTask(foundTask), nil
}

func (t *Task) Create(p *domain_obj.Task) error {
	type iID struct {
		Id int64
	}
	var iid *iID
	t.db.Gdb.Table("importances").Where("name = ?", p.ImportanceName).Find(&iid)

	taskToCreate := p.ToRecord(iid.Id)
	result := t.db.Gdb.Create(taskToCreate)

	return result.Error
}

func (t *Task) Update(p *domain_obj.Task) (*domain_obj.Task, error) {
	type iID struct {
		Id int64
	}
	var iid *iID
	t.db.Gdb.Table("importances").Where("name = ?", p.ImportanceName).Find(&iid)

	taskToUpdate := p.ToRecord(iid.Id)
	result := t.db.Gdb.Updates(taskToUpdate)
	if result.Error != nil {
		return nil, fmt.Errorf("タスク更新処理に失敗しました　%w", result.Error)
	}
	if result.RowsAffected != 1 {
		return nil, fmt.Errorf("タスクは更新されていません")
	}

	// memo: ORMによっては更新後の内容がUpdateメソッドの戻り値で得られる場合もありDB周りの実装をインフラ層に閉じ込めるため、
	// また更新処理後は更新後の内容を返すのがRESTapiでの定石だと思うので、更新後のタスクを取得し直して返却している
	// （pをそのまま返すのもありだと思うが、実際にDBに保存されている内容を返すべきだと思ったので取得し直している）
	updatedTask, err := t.GetById(p.Id)
	if err != nil {
		return nil, fmt.Errorf("更新処理成功後、更新内容取得に失敗しました　%w", err)
	}

	return updatedTask, result.Error
}

func NewTask(db *config.DB) *Task {
	return &Task{db}
}
