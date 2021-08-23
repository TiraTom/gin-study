package domain

import (
	"fmt"

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
	// TODO 多分もっといい割り振り方法がある
	switch i.Name {
	case "MEDIUM":
		return gr.Importance_MEDIUM, nil
	case "HIGH":
		return gr.Importance_HIGH, nil
	case "LOW":
		return gr.Importance_LOW, nil
	case "VERY_HIGH":
		return gr.Importance_VERY_HIGH, nil
	default:
		// Importanceをnilにできないので仮設定。エラーを返すのでそこでエラー検知できる想定。
		return gr.Importance_VERY_HIGH, fmt.Errorf("想定していない重要度ラベル名です %s", i.Name)
	}
}

func NewImportance(ir *infrastructure.ImportanceRecord) *Importance {
	return &Importance{
		Name:  ir.Name,
		Level: ir.Level,
	}
}
