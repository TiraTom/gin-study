package infrastructure

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain"
)

type ImportanceRepository struct {
	db *config.DB
}

func (i *ImportanceRepository) GetAll() *[]domain.Importance {
	result := &[]domain.Importance{}
	i.db.Gdb.Find(&result)

	return result
}

func NewImportanceRepository(db *config.DB) *ImportanceRepository {
	return &ImportanceRepository{db}
}
