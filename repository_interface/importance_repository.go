package repository_interface

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain"
	"github.com/Tiratom/gin-study/infrastructure"
)

type Importance interface {
	GetAll() *[]domain.Importance
}

func NewImportance(db *config.DB) Importance {
	return infrastructure.NewImportanceRepository(db)
}
