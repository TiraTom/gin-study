package infrastructure

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain"
)

type ImportanceImpl struct {
	db *config.DB
}

func (ii *ImportanceImpl) GetAll() *[]domain.Importance {
	result := &[]domain.Importance{}
	ii.db.GormConnect().Find(&result)

	return result
}

func NewImportanceImpl(db *config.DB) *ImportanceImpl {
	return &ImportanceImpl{db}
}
