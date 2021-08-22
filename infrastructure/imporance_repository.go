package infrastructure

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain"
)

type ImportanceRepository interface {
	GetAll() *[]domain.Importance
}

type importanceRepository struct {
	db *config.DB
}

func (i *importanceRepository) GetAll() *[]domain.Importance {
	result := &[]domain.Importance{}
	i.db.Gdb.Find(&result)

	return result
}

func NewImportanceRepository(db *config.DB) *importanceRepository {
	return &importanceRepository{db}
}
