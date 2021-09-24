package domain_obj

import (
	gr "github.com/Tiratom/gin-study/grpc"
	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
)

type Importance struct {
	// 重要度ラベル
	Name string
	// 重要度
	Level int
}

func (i *Importance) ToDto() (gr.Importance, error) {
	return gr.Importance{
		Name:  i.Name,
		Level: uint32(i.Level),
	}, nil
}

func NewImportance(ir *infrastructure.Importance) *Importance {
	return &Importance{
		Name:  ir.Name,
		Level: ir.Level,
	}
}
