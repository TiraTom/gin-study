package domain_obj

import (
	"reflect"
	"testing"

	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
)

func TestNewImportance(t *testing.T) {
	type args struct {
		ir *infrastructure.Importance
	}
	tests := []struct {
		name string
		args args
		want *Importance
	}{
		{
			name: "通常パターン",
			args: args{
				ir: &infrastructure.Importance{
					Id:    1,
					Name:  "HOGE",
					Level: 1,
				},
			},
			want: &Importance{
				Name:  "HOGE",
				Level: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImportance(tt.args.ir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}
