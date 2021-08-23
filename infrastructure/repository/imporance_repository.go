package infrastructure

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
)

type ImportanceRepository struct {
	db *config.DB
}

func (i *ImportanceRepository) GetAll() []*domain.Importance {
	var records []*infrastructure.ImportanceRecord
	i.db.Gdb.Find(&records)

	result := make([]*domain.Importance, len(records))
	for i, v := range records {
		result[i] = domain.NewImportance(v)
	}

	return result
}

func NewImportanceRepository(db *config.DB) *ImportanceRepository {
	return &ImportanceRepository{db}
}
