package domain_obj

import (
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
)

type Importance struct {
	// 重要度ラベル
	Name string
	// 重要度
	Level int
}

func NewImportance(ir *infrastructure.Importance) *Importance {
	return &Importance{
		Name:  ir.Name,
		Level: ir.Level,
	}
}
