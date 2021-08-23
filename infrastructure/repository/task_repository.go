package infrastructure

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
)

type TaskRepository struct {
	db *config.DB
}

func (t *TaskRepository) GetAll() []*domain.Task {
	var records []*infrastructure.TaskAndImportanceRecord
	// TODO この書き方だとimportance系の値を入れてくれない理由がわからない（selectの並び順に決まりがある？）
	// t.db.Gdb.Table("tasks").Select("tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_at as registered_at", "tasks.deadline as deadline", "tasks.updated_at as updated_at", "importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Scan(&records)
	t.db.Gdb.Table("tasks").Select("importances.id as importance_id", "importances.name as importance_name", "importances.level as importance_level", "tasks.id as id", "tasks.name as name", "tasks.details as details", "tasks.registered_at as registered_at", "tasks.deadline as deadline", "tasks.updated_at as updated_at").Joins("LEFT JOIN importances ON tasks.importance_id = importances.id").Scan(&records)

	result := make([]*domain.Task, len(records))
	for i, v := range records {
		result[i] = domain.NewTask(v)
	}

	return result
}

func NewTaskRepository(db *config.DB) *TaskRepository {
	return &TaskRepository{db}
}
