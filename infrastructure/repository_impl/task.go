package repository_impl

import (
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

func (t *Task) Create(p *domain_obj.Task) error {
	type iID struct {
		Id int64
	}
	var i *iID

	t.db.Gdb.Table("importances").Select("importances.id").Where("name = ?", p.ImportanceName).Find(&i)

	taskToCreate := p.ToRecord(i.Id)
	result := t.db.Gdb.Create(taskToCreate)

	return result.Error
}

func NewTask(db *config.DB) *Task {
	return &Task{db}
}
