package repository_impl

import (
	"fmt"

	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/infrastructure/record"
	"gorm.io/gorm"
)

type Task struct {
	db *config.DB
}

func (t *Task) GetAll() (*domain_obj.Tasks, error) {
	var records []*record.TaskAndImportance
	// TODO この書き方だとimportance系の値を入れてくれない理由がわからない（selectの並び順に決まりがある？）
	// t.db.Gdb.Table("tasks").Select("tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_time as registered_time", "tasks.deadline as deadline", "tasks.updated_time as updated_time", "importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Scan(&records)
	result := t.db.Gdb.Table("tasks").Select("importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level", "tasks.version as version", "tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_time as registered_time", "tasks.deadline as deadline", "tasks.updated_time as updated_time").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Scan(&records)

	return domain_obj.NewTasks(records), result.Error
}

func (t *Task) GetById(id string) (*domain_obj.Task, error) {
	var foundTask *record.TaskAndImportance
	result := t.db.Gdb.Table("tasks").Select("importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level", "tasks.version as version", "tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_time as registered_time", "tasks.deadline as deadline", "tasks.updated_time as updated_time").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Where("tasks.id = ?", id).First(&foundTask)

	if result.Error != nil {
		return nil, fmt.Errorf("該当のタスクは存在しません; %w", result.Error)
	}

	return domain_obj.NewTask(foundTask), nil
}

func (t *Task) Create(p *domain_obj.Task) (*domain_obj.Task, error) {
	var iid *domain_obj.ImportanceID
	t.db.Gdb.Table("importances").Where("name = ?", p.ImportanceName).Find(&iid)

	taskToCreate := p.ToRecord(iid.Id)
	result := t.db.Gdb.Create(taskToCreate)

	if result.Error != nil {
		return nil, fmt.Errorf("タスク作成処理に失敗しました(作成したかったTask=%v); %w", p, result.Error)
	}

	createdTask, err := t.GetById(p.Id)
	if err != nil {
		return nil, fmt.Errorf("作成処理成功後、作成内容取得に失敗しました(作成したtask=%v); %w", p, err)
	}

	return createdTask, result.Error
}

func (t *Task) Update(p *domain_obj.Task) (*domain_obj.Task, error) {
	var iid *domain_obj.ImportanceID
	fResult := t.db.Gdb.Table("importances").Where("name = ?", p.ImportanceName).Find(&iid)
	if fResult.Error != nil {
		return nil, fmt.Errorf("重要度ラベルの取得時にエラーが発生しました(importanceName=%v); %w", p.ImportanceName, fResult.Error)
	}
	if !iid.IsValid() {
		return nil, fmt.Errorf("重要度ラベルの取得に失敗しました(importanceName=%v)", p.ImportanceName)
	}

	taskToUpdate := p.ToRecord(iid.Id)
	uResult := t.db.Gdb.Updates(taskToUpdate)
	if uResult.Error != nil {
		return nil, fmt.Errorf("タスク更新処理に失敗しました(更新を試みたタスク=%v); %w", taskToUpdate, uResult.Error)
	}
	if uResult.RowsAffected != 1 {
		return nil, fmt.Errorf("タスクは更新されていません(更新を試みたタスク=%v)", taskToUpdate)
	}

	// memo: ORMによっては更新後の内容がUpdateメソッドの戻り値で得られる場合もありDB周りの実装をインフラ層に閉じ込めるため、
	// また更新処理後は更新後の内容を返すのがRESTapiでの定石だと思うので、更新後のタスクを取得し直して返却している
	// （pをそのまま返すのもありだと思うが、実際にDBに保存されている内容を返すべきだと思ったので取得し直している）
	updatedTask, err := t.GetById(p.Id)
	if err != nil {
		return nil, fmt.Errorf("更新処理成功後、更新内容取得に失敗しました;　%w", err)
	}

	return updatedTask, uResult.Error
}

func (t *Task) Delete(id string) error {
	result := t.db.Gdb.Where("id = ?", id).Delete(&Task{})

	if result.Error != nil {
		return fmt.Errorf("削除対象のTask検索でエラーが発生しました; %w", result.Error)
	}

	if result.RowsAffected != 1 {
		return fmt.Errorf("削除対象のタスク（id=%s）は存在しません", id)
	}

	return nil
}

func (t *Task) Search(p *domain_obj.TaskSearchCondition) (*domain_obj.Tasks, error) {
	var foundTasks []*record.TaskAndImportance

	var result *gorm.DB
	if p.IsDeadlineIncludedInCondition() {
		dcs, err := p.AsDeadlineConditionSentence()
		if err != nil {
			return nil, err
		}
		result = t.db.Gdb.Table("tasks").Select("importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level", "tasks.version as version", "tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_time as registered_time", "tasks.deadline as deadline", "tasks.updated_time as updated_time").Where(p.AsSelectConditionMap()).Where(dcs, p.Deadline).Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Find(&foundTasks)
	} else {
		result = t.db.Gdb.Table("tasks").Select("importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level", "tasks.version as version", "tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_time as registered_time", "tasks.deadline as deadline", "tasks.updated_time as updated_time").Where(p.AsSelectConditionMap()).Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Find(&foundTasks)
	}

	return domain_obj.NewTasks(foundTasks), result.Error
}

func NewTask(db *config.DB) *Task {
	return &Task{db}
}
