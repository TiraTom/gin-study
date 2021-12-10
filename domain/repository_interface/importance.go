//go:generate mockgen -source=importance.go -destination=../../mock/mock_repository_interface/importance.go
package repository_interface

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/infrastructure/repository_impl"
)

type Importance interface {
	GetAll() ([]*domain_obj.Importance, error)
}

func NewImportance(db *config.DB) Importance {
	return repository_impl.NewImportance(db)
}
