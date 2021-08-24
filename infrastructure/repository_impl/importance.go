package repository_impl

import (
	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/infrastructure/record"
)

type Importance struct {
	db *config.DB
}

func (i *Importance) GetAll() []*domain_obj.Importance {
	var records []*record.Importance
	i.db.Gdb.Table("importances").Find(&records)

	result := make([]*domain_obj.Importance, len(records))
	for i, v := range records {
		result[i] = domain_obj.NewImportance(v)
	}

	return result
}

func NewImportance(db *config.DB) *Importance {
	return &Importance{db}
}
