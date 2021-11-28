package repository_impl

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/infrastructure/record"
)

type Importance struct {
	db *config.DB
}

func (i *Importance) GetAll() ([]*domain_obj.Importance, error) {
	var records []*record.Importance
	result := i.db.Gdb.Table("importances").Find(&records)

	importances := make([]*domain_obj.Importance, len(records))
	for i, v := range records {
		importances[i] = domain_obj.NewImportance(v)
	}

	return importances, result.Error
}

func NewImportance(db *config.DB) *Importance {
	return &Importance{db}
}
